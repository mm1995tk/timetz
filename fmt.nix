{pkgs, ...}: {
  projectRootFile = "flake.nix";
  programs.alejandra.enable = true;
  programs.gofmt.enable = true;
  programs.taplo.enable = true;
  programs.prettier.enable = true;
  enableDefaultExcludes = true;
  settings.global.excludes = [
    ".envrc"
  ];
}
