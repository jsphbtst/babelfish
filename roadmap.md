
Babelfish Features Roadmap

[X] GPT4 Translation
[] Google Cloud Translate API (https://pkg.go.dev/cloud.google.com/go/translate/apiv3)
[X] Translation History
[X] Local Cache Translation History
[X] Local Cache Breakdowns History
[X] Translation Breakdown via GPT4
[X] Add Configs Default Language
[X] Add Capability to Change Configs
[X] Make Sure to Check that Target Language Is Legitimate in Configs
[] Add Remote Data Sync Capability

Notes:
- since we're saving records locally, the CLI tool should work without internet
IFF (math notation, bitch) the query conditions are local
- given above, we might have to move internet check conditions to somewhere else
in the code, as opposed to being inside the initialization part of it
- moreover, the file checks and initializations is dogwater, can probs break that
down into a function so more refactors there are absolutely needed
