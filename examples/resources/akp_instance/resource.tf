resource "akp_instance" "example" {
  name = "test"
  argocd = {
    spec = {
      description = "test-inst"
      instance_spec = {
        declarative_management_enabled = false
        backend_ip_allow_list_enabled  = true
        image_updater_enabled          = true
        ip_allow_list = [
          {
            description = "dummy entry2"
            ip          = "1.2.3.4"
          },
        ]
        cluster_customization_defaults = {
          auto_upgrade_disabled = true
        }
        appset_policy = {
          policy          = "create-only"
          override_policy = true
        }
        host_aliases = [
          {
            hostnames = ["test.example.com"]
            ip        = "1.2.3.4"
          },
        ]
        crossplane_extension = {
          resources = [
            {
              group = "*.example.crossplane.*",
            }
          ]
        }
        agent_permissions_rules = [
          {
            api_groups = ["batch"]
            resources  = ["jobs"]
            verbs      = ["create"]
          }
        ]
        fqdn = "test.example.com"
        appset_plugins = [
          {
            # name needs to start with plugin-
            name = "plugin-test"
            # the secret that refers to
            token           = "$application-set-secret:token"
            base_url        = "https://example.com"
            request_timeout = 30
          }
        ]
        app_in_any_namespace_config = {
          enabled = true
        }
      }
      version = "v2.11.4"
    }
  }
  argocd_cm = {
    # When configuring `argocd_cm`, there is generally no need to set all of these keys. If you do not set a key, the API will set suitable default values.
    # Please note that the API will disallow the setting of  any key which isn't a known configuration option in `argocd-cm`.
    #
    # NOTE:
    # `admin.enabled` can be set to `false` to disable the admin login.
    # To enable the admin account, set `accounts.admin: "login"`, and to disable the admin login, set `admin.enabled: false`. They are mutually exclusive.
    "exec.enabled"                   = true
    "ga.anonymizeusers"              = false
    "helm.enabled"                   = true
    "kustomize.enabled"              = true
    "server.rbac.log.enforce.enable" = false
    "statusbadge.enabled"            = false
    "ui.bannerpermanent"             = false
    "users.anonymous.enabled"        = true

    "kustomize.buildOptions" = "--load_restrictor none"
    "accounts.alice"         = "login"
    "dex.config"             = <<-EOF
        connectors:
          # GitHub example
          - type: github
            id: github
            name: GitHub
            config:
              clientID: aabbccddeeff00112233
              clientSecret: $dex.github.clientSecret
              orgs:
              - name: your-github-org
        EOF
  }
  argocd_rbac_cm = {
    "policy.default" = "role:readonly"
    "policy.csv"     = <<-EOF
         p, role:org-admin, applications, *, */*, allow
         p, role:org-admin, clusters, get, *, allow
         g, your-github-org:your-team, role:org-admin
         EOF
  }
  argocd_secret = {
    "dex.github.clientSecret" = "my-github-oidc-secret"
    "webhook.github.secret"   = "shhhh! it'   s a github secret"
  }
  application_set_secret = {
    "my-appset-secret" = "xyz456"
  }
  argocd_notifications_cm = {
    "trigger.on-sync-status-unknown" = <<-EOF
        - when: app.status.sync.status == 'Unknown'
          send: [my-custom-template]
      EOF
    "template.my-custom-template"    = <<-EOF
        message: |
          Application details: {{.context.argocdUrl}}/applications/{{.app.metadata.name}}.
      EOF
    "defaultTriggers"                = <<-EOF
        - on-sync-status-unknown
      EOF
  }
  argocd_notifications_secret = {
    email-username = "test@argoproj.io"
    email-password = "password"
  }
  argocd_image_updater_ssh_config = {
    config = <<-EOF
      Host *
            PubkeyAcceptedAlgorithms +ssh-rsa
            HostkeyAlgorithms +ssh-rsa
            HostkeyAlgorithms2 +ssh-rsa
    EOF
  }
  argocd_image_updater_config = {
    "registries.conf" = <<-EOF
      registries:
        - prefix: docker.io
          name: Docker2
          api_url: https://registry-1.docker.io
          credentials: secret:argocd/argocd-image-updater-secret#my-docker-credentials
    EOF
    "git.email"       = "akuitybot@akuity.io"
    "git.user"        = "akuitybot"
  }
  argocd_image_updater_secret = {
    my-docker-credentials = "abcd1234"
  }
  argocd_ssh_known_hosts_cm = {
    # When configuring the known host list, make sure to add the following default ones before adding your own hosts.
    ssh_known_hosts = <<EOF
[ssh.github.com]:443 ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBEmKSENjQEezOmxkZMy7opKgwFB9nkt5YRrYMjNuG5N87uRgg6CLrbo5wAdT/y6v0mKV0U2w0WZ2YB/++Tpockg=
[ssh.github.com]:443 ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIOMqqnkVzrm0SdG6UOoqKLsabgH5C9okWi0dh2l9GKJl
[ssh.github.com]:443 ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQCj7ndNxQowgcQnjshcLrqPEiiphnt+VTTvDP6mHBL9j1aNUkY4Ue1gvwnGLVlOhGeYrnZaMgRK6+PKCUXaDbC7qtbW8gIkhL7aGCsOr/C56SJMy/BCZfxd1nWzAOxSDPgVsmerOBYfNqltV9/hWCqBywINIR+5dIg6JTJ72pcEpEjcYgXkE2YEFXV1JHnsKgbLWNlhScqb2UmyRkQyytRLtL+38TGxkxCflmO+5Z8CSSNY7GidjMIZ7Q4zMjA2n1nGrlTDkzwDCsw+wqFPGQA179cnfGWOWRVruj16z6XyvxvjJwbz0wQZ75XK5tKSb7FNyeIEs4TT4jk+S4dhPeAUC5y+bDYirYgM4GC7uEnztnZyaVWQ7B381AK4Qdrwt51ZqExKbQpTUNn+EjqoTwvqNj4kqx5QUCI0ThS/YkOxJCXmPUWZbhjpCg56i+2aB6CmK2JGhn57K5mj0MNdBXA4/WnwH6XoPWJzK5Nyu2zB3nAZp+S5hpQs+p1vN1/wsjk=
bitbucket.org ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBPIQmuzMBuKdWeF4+a2sjSSpBK0iqitSQ+5BM9KhpexuGt20JpTVM7u5BDZngncgrqDMbWdxMWWOGtZ9UgbqgZE=
bitbucket.org ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIIazEu89wgQZ4bqs3d63QSMzYVa0MuJ2e2gKTKqu+UUO
bitbucket.org ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDQeJzhupRu0u0cdegZIa8e86EG2qOCsIsD1Xw0xSeiPDlCr7kq97NLmMbpKTX6Esc30NuoqEEHCuc7yWtwp8dI76EEEB1VqY9QJq6vk+aySyboD5QF61I/1WeTwu+deCbgKMGbUijeXhtfbxSxm6JwGrXrhBdofTsbKRUsrN1WoNgUa8uqN1Vx6WAJw1JHPhglEGGHea6QICwJOAr/6mrui/oB7pkaWKHj3z7d1IC4KWLtY47elvjbaTlkN04Kc/5LFEirorGYVbt15kAUlqGM65pk6ZBxtaO3+30LVlORZkxOh+LKL/BvbZ/iRNhItLqNyieoQj/uh/7Iv4uyH/cV/0b4WDSd3DptigWq84lJubb9t/DnZlrJazxyDCulTmKdOR7vs9gMTo+uoIrPSb8ScTtvw65+odKAlBj59dhnVp9zd7QUojOpXlL62Aw56U4oO+FALuevvMjiWeavKhJqlR7i5n9srYcrNV7ttmDw7kf/97P5zauIhxcjX+xHv4M=
github.com ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBEmKSENjQEezOmxkZMy7opKgwFB9nkt5YRrYMjNuG5N87uRgg6CLrbo5wAdT/y6v0mKV0U2w0WZ2YB/++Tpockg=
github.com ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIOMqqnkVzrm0SdG6UOoqKLsabgH5C9okWi0dh2l9GKJl
github.com ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQCj7ndNxQowgcQnjshcLrqPEiiphnt+VTTvDP6mHBL9j1aNUkY4Ue1gvwnGLVlOhGeYrnZaMgRK6+PKCUXaDbC7qtbW8gIkhL7aGCsOr/C56SJMy/BCZfxd1nWzAOxSDPgVsmerOBYfNqltV9/hWCqBywINIR+5dIg6JTJ72pcEpEjcYgXkE2YEFXV1JHnsKgbLWNlhScqb2UmyRkQyytRLtL+38TGxkxCflmO+5Z8CSSNY7GidjMIZ7Q4zMjA2n1nGrlTDkzwDCsw+wqFPGQA179cnfGWOWRVruj16z6XyvxvjJwbz0wQZ75XK5tKSb7FNyeIEs4TT4jk+S4dhPeAUC5y+bDYirYgM4GC7uEnztnZyaVWQ7B381AK4Qdrwt51ZqExKbQpTUNn+EjqoTwvqNj4kqx5QUCI0ThS/YkOxJCXmPUWZbhjpCg56i+2aB6CmK2JGhn57K5mj0MNdBXA4/WnwH6XoPWJzK5Nyu2zB3nAZp+S5hpQs+p1vN1/wsjk=
gitlab.com ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBFSMqzJeV9rUzU4kWitGjeR4PWSa29SPqJ1fVkhtj3Hw9xjLVXVYrU9QlYWrOLXBpQ6KWjbjTDTdDkoohFzgbEY=
gitlab.com ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIAfuCHKVTjquxvt6CM6tdG4SLp1Btn/nOeHHE5UOzRdf
gitlab.com ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCsj2bNKTBSpIYDEGk9KxsGh3mySTRgMtXL583qmBpzeQ+jqCMRgBqB98u3z++J1sKlXHWfM9dyhSevkMwSbhoR8XIq/U0tCNyokEi/ueaBMCvbcTHhO7FcwzY92WK4Yt0aGROY5qX2UKSeOvuP4D6TPqKF1onrSzH9bx9XUf2lEdWT/ia1NEKjunUqu1xOB/StKDHMoX4/OKyIzuS0q/T1zOATthvasJFoPrAjkohTyaDUz2LN5JoH839hViyEG82yB+MjcFV5MU3N1l1QL3cVUCh93xSaua1N85qivl+siMkPGbO5xR/En4iEY6K2XPASUEMaieWVNTRCtJ4S8H+9
ssh.dev.azure.com ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQC7Hr1oTWqNqOlzGJOfGJ4NakVyIzf1rXYd4d7wo6jBlkLvCA4odBlL0mDUyZ0/QUfTTqeu+tm22gOsv+VrVTMk6vwRU75gY/y9ut5Mb3bR5BV58dKXyq9A9UeB5Cakehn5Zgm6x1mKoVyf+FFn26iYqXJRgzIZZcZ5V6hrE0Qg39kZm4az48o0AUbf6Sp4SLdvnuMa2sVNwHBboS7EJkm57XQPVU3/QpyNLHbWDdzwtrlS+ez30S3AdYhLKEOxAG8weOnyrtLJAUen9mTkol8oII1edf7mWWbWVf0nBmly21+nZcmCTISQBtdcyPaEno7fFQMDD26/s0lfKob4Kw8H
vs-ssh.visualstudio.com ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQC7Hr1oTWqNqOlzGJOfGJ4NakVyIzf1rXYd4d7wo6jBlkLvCA4odBlL0mDUyZ0/QUfTTqeu+tm22gOsv+VrVTMk6vwRU75gY/y9ut5Mb3bR5BV58dKXyq9A9UeB5Cakehn5Zgm6x1mKoVyf+FFn26iYqXJRgzIZZcZ5V6hrE0Qg39kZm4az48o0AUbf6Sp4SLdvnuMa2sVNwHBboS7EJkm57XQPVU3/QpyNLHbWDdzwtrlS+ez30S3AdYhLKEOxAG8weOnyrtLJAUen9mTkol8oII1edf7mWWbWVf0nBmly21+nZcmCTISQBtdcyPaEno7fFQMDD26/s0lfKob4Kw8H
      EOF
  }
  argocd_tls_certs_cm = {
    "server.example.com" = <<EOF
          -----BEGIN CERTIFICATE-----
          ......
          -----END CERTIFICATE-----
      EOF
  }
  repo_credential_secrets = {
    repo-my-private-https-repo = {
      url                = "https://github.com/argoproj/argocd-example-apps"
      password           = "my-ppassword"
      username           = "my-username"
      insecure           = true
      forceHttpBasicAuth = true
      enableLfs          = true
    },
    repo-my-private-ssh-repo = {
      url           = "ssh://git@github.com/argoproj/argocd-example-apps"
      sshPrivateKey = <<EOF
      # paste the sshPrivateKey data here
      EOF
      insecure      = true
      enableLfs     = true
    }
  }
  repo_template_credential_secrets = {
    repo-argoproj-https-creds = {
      url      = "https://github.com/argoproj"
      type     = "helm"
      password = "my-password"
      username = "my-username"
    }
  }
  config_management_plugins = {
    "kasane" = {
      image   = "gcr.io/kasaneapp/kasane"
      enabled = true
      spec = {
        init = {
          command = [
            "kasane",
            "update"
          ]
        }
        generate = {
          command = [
            "kasane",
            "show"
          ]
        }
      }
    }
    "tanka" = {
      enabled = true
      image   = "grafana/tanka:0.25.0"
      spec = {
        discover = {
          file_name = "jsonnetfile.json"
        }
        generate = {
          args = [
            "tk show environments/$PARAM_ENV --dangerous-allow-redirect",
          ]
          command = [
            "sh",
            "-c",
          ]
        }
        init = {
          command = [
            "jb",
            "update",
          ]
        }
        parameters = {
          static = [
            {
              name     = "env"
              required = true
              string   = "default"
            },
          ]
        }
        preserve_file_mode = false
        version            = "v1.0"
      }
    },
  }
}
