{
	inputs = {
		nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
	};
	outputs = { nixpkgs, ... }:
	let systems = nixpkgs.lib.platforms.unix;
	in {
		packages = builtins.listToAttrs (builtins.map (system: { name = system; value =
		let pkgs = import nixpkgs { inherit system; };
		in { default = pkgs.buildGoModule {
			pname = "lazynix";
			version = "v0.1.0";
			vendorHash = null;
			src = ./.;
		};};}) systems );
	};
}
