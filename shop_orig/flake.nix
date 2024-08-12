{
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-23.05";
    nixpkgs-unstable.url = "github:NixOS/nixpkgs/nixos-unstable";
    systems.url = "github:nix-systems/default";
    devenv.url = "github:cachix/devenv";
    flake-parts.url = "github:hercules-ci/flake-parts";
    treefmt-nix.url = "github:numtide/treefmt-nix";
  };

  nixConfig = {
    extra-trusted-public-keys = "devenv.cachix.org-1:w1cLUi8dv3hnoSPGAuibQv+f9TZLr6cv/Hm9XgU50cw=";
    extra-substituters = "https://devenv.cachix.org";
  };

  outputs = { self, ... } @ inputs:
    inputs.flake-parts.lib.mkFlake { inherit inputs; } {
      systems = import inputs.systems;
      imports = [
        inputs.treefmt-nix.flakeModule
      ];
      perSystem = { config, self', pkgs, lib, system, ... }:
        let
          pkgs-unstable = import inputs.nixpkgs-unstable {
            inherit system;
            # overlays = [ ];
            # config = { };
          };
        in
        {
          _module.args.pkgs = pkgs-unstable;
          # _module.args.pkgs = import inputs.nixpkgs {
          #   inherit system;
          #   overlays = [
          #     # inputs.foo.overlays.default
          #     (final: prev: {
          #       # TODO: patch stuff from nixpkgs-unstable
          #       go = pkgs-unstable.go;
          #     })
          #   ];
          #   config = { };
          # };
          devShells =
            let
              env = import ./devenv.nix { inherit pkgs inputs; lib = pkgs.lib; };
            in
            {
              default = inputs.devenv.lib.mkShell env;
            };
        };
    };
}
