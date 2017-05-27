#!/bin/bash
signatures=$(cat signatures.yaml)

cat <<EOF
package main
var signatures string =
\`
$signatures
\`
EOF
