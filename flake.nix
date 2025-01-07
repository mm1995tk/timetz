{
  description = "gomod2nix flake";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
    gomod2nix.url = "github:nix-community/gomod2nix";
    gomod2nix.inputs.nixpkgs.follows = "nixpkgs";
    gomod2nix.inputs.flake-utils.follows = "flake-utils";
    treefmt-nix.url = "github:numtide/treefmt-nix";
  };

  outputs = {
    self,
    nixpkgs,
    flake-utils,
    gomod2nix,
    treefmt-nix,
  }:
    flake-utils.lib.eachDefaultSystem (system: let
      pkgs = import nixpkgs {
        inherit system;
        overlays = [gomod2nix.overlays.default];
      };

      treefmtEval = treefmt-nix.lib.evalModule pkgs ./fmt.nix;
    in {
      devShells.default = pkgs.mkShell {
        buildInputs = [
          (pkgs.mkGoEnv {pwd = ./.;})
          pkgs.gomod2nix
          pkgs.golangci-lint
        ];
      };

      devShells.gomod2nix = pkgs.mkShell {
        buildInputs = [
          pkgs.gomod2nix
        ];
      };

      formatter = treefmtEval.config.build.wrapper;
      checks.formatting = treefmtEval.config.build.check self;
    });
}
