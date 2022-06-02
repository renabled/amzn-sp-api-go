#! /bin/bash

rm -rf ./selling-partner-api-models
git clone https://github.com/amzn/selling-partner-api-models.git

checksum="$(find ./selling-partner-api-models/models -type f -print0 | sort -z | xargs -0 sha1sum | sha1sum)"

if [ "$(cat ./script/checksum)" == "$checksum" ]; then
  echo "✅ No changes detected"
else
  echo "$checksum" > ./script/checksum

  rm -rf ./api
  mkdir -p ./api

  for f in ./selling-partner-api-models/models/**/*.json; do \
    echo "⚙️ Processing $f..."; \

    if [ "$(basename -s .json $f)" == "catalogItems_2020-12-01" ] || \
        [ "$(basename -s .json $f)" == "listingsItems_2021-08-01" ]; then \
      sed 's/"default": "summaries"/"default": ["summaries"]/g' $f > ./selling-partner-api-models/tmp.json; \
      mkdir -p ./api/"$(basename -s .json $f)"; \
      swagger --quiet generate client -f ./selling-partner-api-models/tmp.json -t ./api/"$(basename -s .json $f)" -c "$(basename -s .json $f)-client" -A "$(basename -s .json $f)" -m "$(basename -s .json $f)-models"; \
      continue; \
    fi; \

    mkdir -p ./api/"$(basename -s .json $f)"; \
    swagger --quiet generate client -f "$f" -t ./api/"$(basename -s .json $f)" -c "$(basename -s .json $f)-client" -A "$(basename -s .json $f)" -m "$(basename -s .json $f)-models"; \
  done

  echo "✅ Done generating!"
fi

rm -rf ./selling-partner-api-models
