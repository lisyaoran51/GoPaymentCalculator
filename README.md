# GoPaymentCalculator


### Usage

use below syntax to compose payment calculator
```
NewCalculator()
    .CompositeBy(<parent calculator>)
    .CompositeBy(<parent calculator>)
    .CompositeBy(<parent calculator>)

```
then run `Calculate(member, point, coin)` to finish your payment calculation.

the calculate flow would be `child >> parent` 

<br>

### Structure

> P.S. the inheritance in golang is implemented by `compositon over inheritance` technique

![](./doc/schema.png?raw=true)

<br>

### Compositie Payment Solution

It uses `composite pattern` to solve composite payment problem through composing different calculators.

<br>

### Scenario

1. #### normal payment situation
    please check out [example_test.go](./example_test.go) `TestEmptyPayment`

2. #### VIP discount situation
    please check out [example_test.go](./example_test.go) `TestDiscount`

3. #### point redeem situation
    please check out [example_test.go](./example_test.go) `TestRedeem`

4. #### 10% discount if point redeem is bigger than 100
    please check out [example_test.go](./example_test.go) `TestDiscountRedeem`

Feel free to ask me if there's any question!

