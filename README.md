# arr-interfaces

Utilities to be used alongside **bitmagnet**, **sonarr** and **radarr**

- TMDB proxy.  
   - Reduce number of hits that **bitmagnet** makes on TMDB.  Caches content that is in **sonarr** or **radarr** and frequently encountered search terms from **bitmagnet** classifier
   - more sophisticated search of known content using **meilisearch**  Enables use of synonyms and stop words in search
   - more sophisticated similarity matching.  Uses [go-edlib](https://github.com/hbollon/go-edlib). Blends levenshstein, LCS, Jaccard, OSA, SorensenDice and Qgram.  
- **bitmagnet** classifier flag file generation
- webhook triggers to regenerate **bitmagnet** classifier flag files on series add in **sonarr** or movie add in **radarr**
- scheduled regeneration of **bitmagnet** classifier flag files

Using the proxy,  load is reduced on **postgres** database and performance is similar to *localsearch* in **bitmagnet**

# Sample *bitmagnet* configuration
```
tmdb:
  api_key: ##your TMDB API key##
  base_url: http://arr-interfaces:3335/tmdb/3
  rate_limit: 1ns
  rate_limit_burst: 1000

classifier:
  workflow: tmdbproxy
```

# Sample *bitmagnet* classifier
- Until [PR #379](https://github.com/bitmagnet-io/bitmagnet/pull/379) is merged **bitmagnet** supports a single supplementary `classifier.yml` file.   Flag files are created in location defined by `XDG_CONFIG_HOME` 
- enhancement to **bitmagnet** has been coded that automatically picks up changed `classifier*.yml` files.  In interim to pick these up, it is necessary to restart **bitmagnet**
- `tmdb_similarity_required` requires change to **bitmagnet**.  Similarity by default is duplicated in **bitmagnet** with less sophisticated approach resulting in good matches being discarded (about 2-3% of good results discarded).  Ready to raise PR.
```
flags:
  local_search_enabled: false
  # dependent on change to bitmagnet.  PR not raised.
  # tmdb_similarity_required: false

# extend the default workflow with a custom workflow to tag torrents containing interesting documents:
workflows:

  tagandkeep:
    # tag active content.  slim down other torrents with few seeders
    - find_match:
      - if_else:
          condition:
            and:
              - "result.contentType == contentType.tv_show"
              - "result.contentSource == 'tmdb' "
              - "result.contentId in flags.sonarr "
              - "result.episodes.filter(e, result.contentId + '_' + e in flags.sonarr_active).size() > 0"
          if_action:
            add_tag: active
          else_action: unmatched

      - if_else:
          condition:
            and:
              - "result.contentType == contentType.tv_show"
              - "result.contentSource == 'tmdb' "
              - "result.contentId in flags.sonarr "
          if_action:
            add_tag: sonarr
          else_action: unmatched

      - if_else:
          condition:
            and:
              - "result.contentType == contentType.movie"
              - "result.contentSource == 'tmdb' "
              - "result.contentId in flags.radarr "
          if_action:
            add_tag: radarr
          else_action: unmatched
      - if_else:
          condition: "torrent.seeders <= 5"
          if_action: delete
          else_action: unmatched


  tmdbproxy:
    - run_workflow: default
    - run_workflow: tagandkeep

```

# Sample *arr-interfaces* config
**arr-interfaces** uses same configuration framework as **bitmagnet**.  To view configuration and options
```
$ docker exec -it arr-interfaces sh
$ arr-interfaces config show
```
## Sample

```
quartz:
  test_run: false
  schedule:
    - name: SonarrClassifier
      cron_expr: "0 0 3 * * *"
    - name: RadarrClassifier
      cron_expr: "0 0 3 1,4 * *"

meiliclient:
  master_key: meilisearch
  uri: http://meilisearch:7700
  
tmdbproxy:
  api_key: ## your TMDB API key ##
  compress_cache: true
  sonarr:
    api_key: ## your sonarr API key ##
    url: http://sonarr.home
  radarr:
    api_key: ## your radarr API key ##
    url: http://radarr.home
  meilisearch:
    synonyms:
      au:
        - "australian"
      uk:
        - "gb"
      hells:
        - "hell's"
  tmdb:
    fetch_missing: true

```

# Addition notes
- The shared sample **bitmagnet** `classifier.yml` is designed to keep content that is in self-hosted **sonarr** or **radarr**   This keeps **postgres** database smaller,  which maybe useful for some.  
- The tagging can be used by `torznab` interface of **bitmagnet** when [PR 371](https://github.com/bitmagnet-io/bitmagnet/pull/371) is merged.  With this enabled,  content is picked up more reliably and quickly by **sonarr** and **radarr**
- Intention is to release a fork of **bitmagnet** that has changes that **arr-interfaces** benefits from to for increased reliability and efficieny of using **bitmagnet** as sole `torznab` RSS interface to **sonarr** and **radarr**.   Understandably rate at which PRs can be reviewed / merged is limited on a high quality single maintainer project.
- webhooks need to be configured in **sonarr** amd **radarr** to send events to **arr-interfaces**.  Sample URL `http://192.168.1.20:3335/webhook`