{ inputs, pkgs, lib, ... }:
{
  inherit pkgs inputs;
  modules = [
    {
      pre-commit.hooks = { };
      languages.go.enable = true;
      languages.go.package = pkgs.go;

      # XXX: fish shell is not supported by devenv's enterShell
      # enterShell = ''
      # '';
      packages = lib.optionals
        pkgs.stdenv.isDarwin
        (with pkgs.darwin.apple_sdk; [
        ]) ++
      [
        pkgs.golangci-lint
        pkgs.go-ethereum
      ];
    }
  ];
}
