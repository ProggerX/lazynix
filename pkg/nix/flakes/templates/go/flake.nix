{
	description = "Example Go flake template";

	inputs = {
		nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
	};

	outputs = { nixpkgs, ... }:
	let system = "x86_64-linux";
	pkgs = nixpkgs.legacyPackages.${system};
	in {
		packages.${system}.default = pkgs.buildGoModule {
			pname = "example";
			version = "1.0.0";
			src = ./.;
			vendorHash = null; # If you have external dependencies, use 'vendor' directory or specify this hash
		};
		devShells.${system}.default = pkgs.mkShell {
			name = "example-shell";
			packages = with pkgs; [ go ];
		};
	};
}
