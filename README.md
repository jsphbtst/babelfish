# Babelfish

Babelfish is a CLI tool made by me, Joseph. I'm currently learning some languages right now and I find myself opening and closing ChatGPT and Google Translate every now and then—it's tiresome. Given that I spend a lot of time on my terminal, it only made sense that I somehow transfer that common language learning workflow from there to here.

This tool is only really for myself, so don't expect me to entertain any PRs or whatnot. I'm making this available to the public because it sounds like a fun idea–hello, recruiters!

I still have a lot of work to do that I haven't even put in the roadmap yet—the deeper I get into my language learning journey, the more insights I gather that cause me to rethink learning strategies. Those should soon be reflected as features.

Stay tuned!

## Usage

There are a bunch of features I included to make this work for me (really, idc about you). These are:

1. translate from a source language to a target languages (`babelfish translate`);
2. break down a sentence and offer an explaination (`babelfish breakdown`);
3. view your past translations (`babelfish history list`); and
4. set some defaults (`babelfish configs`).

Admittedly, 4) needs more work on my side since for now, it's only the default target language that's actually useful. In due time, I guess.

### Translate

This feature allos you to translate a sentence from one language to another. I cap it at 180 CWS because I don't want to translate text corpuses. The purpose is really to just translate short sentences—hey, that's what works for me. By default, it uses your config's default target language, but this can be configurable.

The command below translates "what's, up doc?" to your default target language:
```
babelfish translate "what's up, doc?"
```

If you want to explicitly state what language you want to translate to, you can do that too:
```
babelfish translate "what's up, doc?" -t Nihongo
```

### Breakdown

Sometimes you get a translation and you don't know why the hell it's structured the way that it is. For example, I'm native in English and Filipino, so languages like Castellano make sense to me grammatically, since they all roughly follow the Subject-Verb-Object (SVO) pattern. However, languages like Japanese—and sometimes German, though not consistently—follow SOV.

Let's start with the English sentence "The boy eats apples." We're starting with English because lots of us speak it.
```
Subject: The boy
Verb: eats
Object: apples
```

In Filipino, that's "Ang bata ay kumakain ng mansanas":
```
Subject: Ang bata
Verb: kumakain
Object: mansanas
```

Y si, tambien en Castellano, que es "El nino come manzanas":
```
Subject: El nino
Verb: come
Object: manzanas
```

...in Japanese? Well, 男の子がりんごを食べます (Romaji: Otoko no ko ga ringo wo tabemasu)
```
Subject: Otoko no ko (the boy)
Object: ringo (apple)
Verb: tabemasu (eats)
```

Sorry for the nerd sesh and slight flex. But breakdowns are useful for obtaining a deeper understanding, even just with simple sentences. Here's a sample output from a past command I ran:
```
babelfish breakdown "hemos hablado espanyol en casa"

Breakdown: "Hemos hablado español en casa" is a Spanish phrase that translates to "We have spoken Spanish at home" in English. Here's the breakdown:

- "Hemos" is the first-person plural form of the auxiliary verb "haber" in the present perfect tense, which translates to "we have".
- "Hablado" is the past participle form of the verb "hablar," which means "to speak". When used with "hemos", it forms a compound verb tense known as the present perfect, translating to "have spoken".
- "Español" is a noun that translates to "Spanish", the language.
- "En" is a preposition which translates to "in" or "at".
- "Casa" is a noun that translates to "home".

Therefore, each word translated individually would be "we have spoken Spanish in/at home".
```

### History

TODO

### Configs

TODO