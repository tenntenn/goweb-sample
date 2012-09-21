{{define "common/settings"}}
define('common/settings',
    [
    ],
    ()->
        $$ =
            host : "{{.Host}}"
)
{{end}}
