{
  description = "A very basic flake";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
  };

  outputs = {
    nixpkgs,
    flake-utils,
    ...
  }:
    flake-utils.lib.eachDefaultSystem (system: let
      pkgs = nixpkgs.legacyPackages.${system};
    in {
      formatter = pkgs.alejandra;
      devShells = {
        default = pkgs.mkShell {
          packages = with pkgs; [
            bundler
            go_1_22
            go-task
            golangci-lint
            jekyll
            rover
            goreleaser

          ];
        };
      };
    });

}
