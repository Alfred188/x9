	         ________
	___  ___/   __   \
	\  \/  /\____    /
	 >    <    /    / 
	/__/\__\  /____/  

This app will create new links by adding your injection to it.

Sample Use:
```
ls *.x9 | while read file; do x9 -ul $file -c 40 -v "<b/minj,'minj'" -gs all -vs suffix -w ~/wordlist/top25  | nuclei -t ~/nuclei-custom-template/Reflection_discovery.yaml  -silent ; done > output

```

Sample Nuclei code:
```id: parameter-discovery

info:
  name: Headless Parameter Discovery
  author: Hades
  severity: info
  description: Takes parameters fuzz for the reflection
  tags: headless

http:
  - method: GET
    path:
      - "{{BaseURL}}"

    # headers contain the headers for the request
    headers:
      # Custom user-agent header
      User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/112.0
      # Custom referrer header
      Referer: "{{RootURL}}"
      # Custom request origin
      Origin: "{{RootURL}}"
      # Custom cookie
      Cookie: pc_auth=zpx03hxz8avb0o9l0765i2aqmwmux9pf3aod8duh73zu3846f92dsiyehymtv0ip

    matchers-condition: or
    matchers:
      - type: word
        words:
          - <b>minj
        part: body
      - type: word
        words:
          - <b/minj
        part: body
      - type: word
        words:
          - '"minj""'
        part: body
      - type: word
        words:
          - "'minj''"
        part: body
    ```