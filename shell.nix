with import <nixpkgs> { };
with pkgs;
stdenv.mkDerivation {
  name = "elvish-plugins";
  nativeBuildInputs = [
    pkg-config
    gtk3.dev
  ];
}
