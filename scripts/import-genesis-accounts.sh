for row in $(cat ./accounts.json | jq -r '.[] | @base64'); do
    _jq() {
        echo ${row} | base64 --decode | jq -r ${1}
    }

    meled add-genesis-account $(_jq '.address') $(_jq '.umelg')umelg,$(_jq '.umelc')umelc
    echo "$(_jq '.address') imported."
done