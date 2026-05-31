{
  lib,
  pkgs,
  config,
  ...
}:
let
  inherit (lib)
    mkIf
    mkEnableOption
    mkOption
    types
    ;
  cfg = config.services.fleur;
in
{
  options.services.fleur = {
    enable = mkEnableOption "Fleur service";

    package = mkOption {
      type = types.package;
      default = pkgs.callPackage ../. { };
      description = ''
        The package to use.
      '';
    };

    openPort = mkOption {
      type = types.bool;
      default = false;
      description = ''
        Whether to open ports for each server.
      '';
    };

    user = mkOption {
      type = types.str;
      default = "fleur";
      description = ''
        User under which the fleur server will run.
      '';
    };

    group = mkOption {
      type = types.str;
      default = "fleur";
      description = ''
        Group under which the fleur server will run.
      '';
    };

    directory = mkOption {
      type = types.nullOr types.str;
      default = null;
      description = ''
        Directory path that will be the root of the Gopher server.
      '';
    };

    port = mkOption {
      type = types.nullOr types.int;
      default = null;
      description = ''
        Gopher server port
      '';
    };

    verbose = mkOption {
      type = types.nullOr types.bool;
      default = false;
      description = ''
        Enable verbose logs
      '';
    };

    domain = mkOption {
      type = types.nullOr types.str;
      default = null;
      description = ''
        Gopher domain.
      '';
    };

    autoInlineText = mkOption {
      type = types.nullOr types.bool;
      default = false;
      description = ''
        Relax non compliant text error and convert to gophermap inline text.
      '';
    };

    personalGopherspaces = mkOption {
      type = types.nullOr types.bool;
      default = false;
      description = ''
        Relax non compliant text error and convert to gophermap inline text.
      '';
    };
  };

  config = mkIf cfg.enable {
    users = {
      users = {
        ${cfg.user} = {
          description = "fleur service user";
          isSystemUser = true;
          inherit (cfg) group;
        };
      };
      groups.${cfg.group} = { };
    };

    systemd.services.fleur = {
      description = "Fleur Server";
      wantedBy = [ "multi-user.target" ];
      after = [ "network.target" ];

      serviceConfig = {
        ExecStart =
          "${lib.getExe cfg.package}"
          + (lib.optionalString (cfg.directory != null) " -directory ${cfg.directory}")
          + (lib.optionalString (cfg.domain != null) " -domain ${cfg.domain}")
          + (lib.optionalString (cfg.port != null) " -port ${builtins.toString cfg.port}")
          + (lib.optionalString cfg.verbose " -verbose")
          + (lib.optionalString cfg.autoInlineText " -enable-auto-inline-text")
          + (lib.optionalString cfg.personalGopherspaces " -enable-personal-gopherspaces");

        User = cfg.user;
        Group = cfg.group;

        # Hardening
        PrivateDevices = true;
        PrivateUsers = true;
        ProtectHome = true;
        ProtectKernelLogs = true;
        ProtectKernelModules = true;
        ProtectKernelTunables = true;
        RestrictAddressFamilies = [
          "AF_INET"
          "AF_INET6"
        ];
        RestrictNamespaces = true;
        SystemCallArchitectures = "native";
        CapabilityBoundingSet = [ "" ];
        DeviceAllow = [ "" ];
        LockPersonality = true;
        PrivateTmp = true;
        ProtectClock = true;
        ProtectControlGroups = true;
        ProtectHostname = true;
        ProtectProc = "invisible";
        RestrictRealtime = true;
        RestrictSUIDSGID = true;
        UMask = "0007";
      };
    };

    networking.firewall = {
      allowedTCPPorts = lib.optional cfg.openPort cfg.port;
    };
  };
}
