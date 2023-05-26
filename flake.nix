{
  description = "Go example flake for Zero to Nix";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs";
    flake-utils.url = "github:numtide/flake-utils";
    gitignore = {
      url = "github:hercules-ci/gitignore.nix";
      inputs.nixpkgs.follows = "nixpkgs";
    };
  };

  outputs = { self, nixpkgs, flake-utils, gitignore }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { inherit system ; };
      in
      {
        defaultPackage = pkgs.buildGoModule {
          name = "devxstats";
          src = gitignore.lib.gitignoreSource ./.;
          vendorSha256 = "sha256-0D+TS/YZIwQasgv46Jl1mUBzwpTteAV6iqBHVDlG1j0=";
        };

        devShells.default = pkgs.mkShellNoCC {
          packages = with pkgs; [
            go_1_20
            gotools
            gopls
          ];
        };
      }
    );
}
