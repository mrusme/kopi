# Kopi

> From Malay _kopi_ (“coffee”) and Hokkien 咖啡 (ko-pi); the latter is also
> derived from the former.\
> -- [Wiktionary](https://en.wiktionary.org/wiki/kopi)

_Kopi_ is a command-line (CLI) coffee journal (or _habit tracker_), designed for
coffee enthusiasts. It lets you track coffee beans, equipment usage, brewing
methods, and individual cups.

The tracked data offers insights into bean and roast preferences, caffeine and
dairy consumption, and equipment usage patterns. This helps users refine their
coffee choices while managing caffeine intake for a more enjoyable and
responsible experience.

_Kopi_ focuses on meticulously tracking every step of the coffee preparation
process. Unlike standard coffee or caffeine tracking apps, it requires more
detailed input to function effectively. If you only want to track the cups of
coffee you drink without monitoring the beans/roasts and the preparation methods
used, _Kopi_ may not be the best fit for you. Think of _Kopi_ as a coffee
journal that lets you rate individual roasts and analyze how different roasts,
brewing methods, and drink types influence your preferences.

_Kopi_ stores everything **locally** and does not send any data to third-party
services unless explicitly mentioned before an operation (e.g. requesting
current exchange rates from the ECB, for price input/conversions of individual
coffee bags). _Kopi_ does not employ analytics libraries/services.

## Workflow

_Kopi_ follows a structured three-step workflow for tracking coffee consumption,
essential for its functionality and insights. The first step is typically done
once—during setup or when equipment changes. The second step is performed
regularly, while the third step, tracking actual consumption, is done for every
cup.

### 1: Add equipment

Coffee equipment is the first piece of information that _Kopi_ requires. Most
coffee enthusiasts will usually have at least two pieces of coffee equipment at
their disposal:

- A coffee grinder
- A coffee maker

This equipment needs to be added to _Kopi_ before tracking coffee brewing and
consumption. To do so, the following command is used:

```sh
▲ kopi equipment add
```

The command will guide the user through the adding process. It is also possible
to provide all required fields as flags to the command. For more information on
how to do so, check the output of `kopi equipment add --help`.

**Note:** If you don’t have coffee brewing equipment and typically get your
coffee from a barista, add their brewing method as a _dummy_ entry. For
espresso-based drinks, add an _espresso maker_; for pourovers or AeroPress, add
a _coffee maker_. However, keep in mind that _Kopi_ is **not** designed for
tracking individual cups from various coffee shops—there are better tools for
that. Instead, _Kopi_ is aimed at coffee enthusiasts who want to document their
coffee experiences, ideally using their own brewing methods.

### 2: Open a bag of coffee

Coffee bags are the second key element in tracking brewing and consumption.
Every cup requires a designated bag of coffee. Previously _opened_ bags can be
used until their beans are depleted, at which point a new bag must be _opened_.
Multiple bags can be _open_ at once, but only one can be selected per cup. For
blends, a dedicated bag must be prepared and then _opened_.

To _open_ a new bag of coffee the following command is used:

```sh
▲ kopi bag open
```

The command will guide the user through the opening process. It is also possible
to provide all required fields as flags to the command. For more information on
how to do so, check the output of `kopi bag open --help`.

To list all open bags, the following command can be used:

```sh
▲ kopi bags list
```

_Note:_ `bags` is an alias of the `bag` command that helps with making the
command read more natural. The command `kopi bag list` works and can as well be
used.

### 3: Drink a cup of coffee

To track a cup of coffee, the following command is used:

```sh
▲ kopi cup drink
```

The command will guide the user through the tracking process. It is also
possible to provide all required fields as flags to the command. For more
information on how to do so, check the output of `kopi cup drink --help`.
