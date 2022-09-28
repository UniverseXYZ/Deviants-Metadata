# Deviants Metadata
## Overview
See general overview & contracts here: [Deviants-Contracts](https://github.com/UniverseXYZ/Deviants-Contracts/)
## Genes interpretation
Each gene is 2 numbers read from right to left. Interpret follows:

- Gene 1 [0:2] - background attribute
- Gene 2 [2:4] - body attribute
- Gene 3 [4:6] - footwear attribute
- Gene 4 [6:8] - left hand attribute
- Gene 5 [8:10] - pants attribute
- Gene 6 [10:12] - torso attribute
- Gene 7 [12:14] - righthand attribute
- Gene 8 [14:16] - head attribute (should be same as the body!)
- Gene 9 [16:18] - eyewear weapon attribute
- Gene 10 [18:20] - headwear attribute
- Gene 11 [20:22] - music attribute

### Note 
- `some traits have special rules`
  - some `left hand` traits must be above `torso` traits
    1. Bag of Herbs
    2. Cheese
    3. Chocolate Bar
    4. Dark Matter Beam
    5. Ice Cream Cone
    6. Peanut Butter
    7. Lizard
    8. Cat
  - one `eyewear` trait must be above `headwear` Traits
    1. Scuba Gear
- Above rules are applied by swapping `left hand` and `torso` and/or `eyewear` with `headwear` order when layering the image 
if a `left hand`/`eyewear` trait corresponds to the ones described
## Deployment
```bash
gcloud functions deploy images-function-test --entry-point TokenMetadata --runtime go116 --trigger-http --allow-unauthenticated
```

## Generate Contract
```bash
abigen --abi app/artifacts/Deviants.abi --pkg contracts --type Deviants --out app/contracts/Deviants.go
```