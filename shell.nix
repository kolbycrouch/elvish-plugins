with import <nixpkgs> { };
with pkgs;
stdenv.mkDerivation {
  name = "elvish-plugins";
  nativeBuildInputs = [
    go
  ];
}
