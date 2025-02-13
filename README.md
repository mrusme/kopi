# Kopi

> From Malay kopi (“coffee”) and Hokkien 咖啡 (ko-pi); the latter is also
> derived from the former.\
> -- [Wiktionary](https://en.wiktionary.org/wiki/kopi)

_Kopi_ is a command-line (CLI) habit tracker specifically designed for tracking
coffee brewing and consumption. It allows tracking coffees (as in bags of
beans), coffee equipment use, as well as individual cups of coffee.

The tracked data gives insights into caffeine and dairy intake, as well as usage
pattern across different equipment and preferences in terms of coffee, to help
users on their way to responsible and joyful consumption of coffee and minimize
side-effects of caffeine.

Kopi stores everything **locally** and does not send any data to third-party
services unless explicitly mentioned before an operation (e.g. requesting
current exchange rates from the ECB, for price input/conversions of individual
coffee bags). **Personal data will never leave your computer, ever.** Kopi does
not employ analytics libraries/services.

## Workflow

Kopi consists of a pre-defined three-step workflow for tracking coffee
consumption. This workflow must be followed in order for Kopi to function and
provide valuable insights. The first step is usually only performed the first
time Kopi is used, or when equipment changes. The second step is performed
regularly. The third step is tracking the actual consumption and is the only
step that is being performed for every consumed cup.

### 1: Add equipment

Coffee equipment is the first piece of information that Kopi requires. Most
coffee enthusiasts will usually have two pieces of coffee equipment at their
disposal:

- A coffee grinder
- A coffee maker

This equipment needs to be added to Kopi before tracking coffee brewing and
consumption. To do so, the following command is used:

```sh
▲ kopi equipment add
```

The command will guide the user through the adding process. It is also possible
to provide all required fields as flags to the command. For more information on
how to do so, check the output of `kopi equipment add --help`.

_Note:_ If you don't happen to have any coffee brewing equipment at your
disposal and you usually get your coffee pre-made by a barista, add the brewing
equipment they use as a _dummy_ entry. If you usually order espresso-based
drinks, you want to add an _espresso maker_. If you order pourovers, or
AeroPress coffee, you want to add a _coffee maker_.

### 2: Open a bag of coffee

Coffee bags are the second information needed for tracking coffee brewing and
consumption. Whenever a cup of coffee is to be enjoyed, a bag of coffee is
required. Previously _opened_ bags of coffee can be used until the coffee beans
in that bag were consumed, at which point a new bag of coffee must be _opened_.
Multiple bags of coffee can be _open_ simultaneously, but only a single bag of
coffee can be chosen per cup of coffee. For coffee blends, the blend must be
prepared in a dedicated bag, which is then _opened_.

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
