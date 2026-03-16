# Powerlevel10k segment for target-cli
#
# Displays active target-cli context profiles in the prompt.
# Each tool with an active profile shows as "tool:profile".
#
# Setup:
#   1. Source this file in your ~/.zshrc (after sourcing powerlevel10k):
#        source /path/to/p10k-target.zsh
#
#   2. Add 'target' to your prompt elements in ~/.p10k.zsh:
#        typeset -g POWERLEVEL9K_LEFT_PROMPT_ELEMENTS=(... target ...)
#      or on the right:
#        typeset -g POWERLEVEL9K_RIGHT_PROMPT_ELEMENTS=(... target ...)
#
#   3. Optional — customise colours and icons in ~/.p10k.zsh:
#        typeset -g POWERLEVEL9K_TARGET_FOREGROUND=208
#        typeset -g POWERLEVEL9K_TARGET_BACKGROUND=236
#        typeset -g POWERLEVEL9K_TARGET_VISUAL_IDENTIFIER_EXPANSION=''

function prompt_target() {
  local parts=()

  [[ -n $TARGET_VAULT_PROFILE     ]] && parts+=("vault:${TARGET_VAULT_PROFILE}")
  [[ -n $TARGET_NOMAD_PROFILE     ]] && parts+=("nomad:${TARGET_NOMAD_PROFILE}")
  [[ -n $TARGET_CONSUL_PROFILE    ]] && parts+=("consul:${TARGET_CONSUL_PROFILE}")
  [[ -n $TARGET_BOUNDARY_PROFILE  ]] && parts+=("boundary:${TARGET_BOUNDARY_PROFILE}")
  [[ -n $TARGET_TERRAFORM_PROFILE ]] && parts+=("tf:${TARGET_TERRAFORM_PROFILE}")

  (( ${#parts[@]} == 0 )) && return

  p10k segment -f "${POWERLEVEL9K_TARGET_FOREGROUND:-208}" \
               -i "${POWERLEVEL9K_TARGET_VISUAL_IDENTIFIER_EXPANSION-}" \
               -t "${(j: :)parts}"
}

# Required for instant prompt compatibility — this segment reads only env vars,
# so it is safe to run during the instant prompt phase.
function instant_prompt_target() {
  prompt_target
}
