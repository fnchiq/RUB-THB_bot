# Tg Bot Road Map

## Текущая версия

Текущая версия бота. Бот используется администратором обменника как внутренний инструмент.

Возможности:

- Курсы биржи Binance
- Курсы биржи Bybit
- Основной процесс:
  1. `/start` Начало
  2. `Главное меню`
  3. Выбор биржи `Binance`, `Bybit`
  4. Выбор фиатной валюты `RUB`, `THB`
  5. Выбор банка `Сбербанк`, `Тинькофф`
  6. Ввод суммы
  7. Получение курса

Текущие лимиты:

-

## Этап 1. Добавление REST API для выполнения аналогичных команд

### REST API сервис

- В сервисе должен быть реализован CORS для локально запущенного сайта `[http://localhost:4000](http://localhost:3000)` что бы к нему мог обращаться веб-фронтенд
- Сервис должен запуститься на порту 5000
- В случае возникновения ошибок, в ответах слдуе возвращать сообщение об ошибке:
  - `{"result": false, "error": "Error message"}`
  - HTTP Status Code 500 для любых ошибок
  - Отдельный код 404 для объектов которые не были найдены
  - В случае возникновения ошибки, так-же следует записать в консоль сообщение: `[ERROR] Programm Module. Error message`, где Programm Module - название модуля в котором произошла ошибка
- Так-же следует логировать в консоль сведения о вхощих запросах, пример: `[INFO] Programm Module. Request: POST /api/GetUsdtFiatRate`

### Метод для получения курса

- Метод должен расчитать курсы для выбранных бирж и выбранных банков для покупки USDT для указанной суммы в фиате.
- Так как на биржах наименование имен банков может отличаться, тут следует выбрать для них специальный алиас и далее в каждой бирже приводить к действительному имени банка
- Метод должен отсортировать ответ по выгоде обмена, от меньшего к большему по полю price

Пример запроса:

Расчет стоимости покупки USDT на биржах Binance, Bybit для фиатной валюты RUB на сумму 1234 рублей

```jsx
POST /api/GetUsdtFiatRate
{
    "exchange": ["Binance", "Bybit"]
    "fiat": "RUB",
    "bank": ["Sber", "Tinkoff"],
    "tradeType": "BUY",
    "fiatAmount": 1234
}
```

Спека ответа:

```json
{
  "result": true,
  "rates": [
    { "exchange": "Binance", "bank": "Sber", "price": 13.61 },
    { "exchange": "Binance", "bank": "Tinkoff", "price": 13.58 },
    { "exchange": "Bybit", "bank": "Sber", "price": 13.13 },
    { "exchange": "Bybit", "bank": "Tinkoff", "price": 13.45 }
  ]
}
```

## Этап 2. Добавление простого веб интерфейса

Пример для Chat GPT запроса

```text
write single html page  with svelte (use svelte linked  from script with public CDN)

Page must contains form with:
* Group "Exchanges", input with type checkbox labeled "Binance",
* Group "Exchanges", input with type checkbox labeled "ByBit"
* Group "Banks", input with type checkbox labeled "Sber"
* Group "Banks", input with type checkbox labeled "Tinkoff"
* Group "Parameters", input with  number type labeled "Fiat Ammount"
* Group "Actions". Submit button
* Group "Actions". Reset button

When pressed Submit button, js code must call POST http://localhost:5000/api/GetUsdtFiatRate with json payload like:

{
    "exchange": ["Binance", "Bybit"]
    "fiat": "RUB",
    "bank": ["Sber", "Tinkoff"],
    "tradeType": "BUY",
    "fiatAmount": 1234
}
response for this request will be JSON:
{
  "result": true,
  "rates": [
    { "exchange": "Binance", "bank": "Sber", "price": 13.61 },
    { "exchange": "Binance", "bank": "Tinkoff", "price": 13.58 },
    { "exchange": "Bybit", "bank": "Sber", "price": 13.13 },
    { "exchange": "Bybit", "bank": "Tinkoff", "price": 13.45 }
  ]
}

this data should be displayed after form in simple table
```

Пример фронтенда:

```html
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <script
      src="https://cdn.jsdelivr.net/npm/svelte@3.42.4/dist/index.mjs"
      defer
    ></script>
    <title>Svelte Form</title>
  </head>
  <body>
    <div id="app"></div>

    <script>
      const app = new App({
        target: document.getElementById("app"),
      });
    </script>
  </body>
</html>

<script>
  import { onMount } from "svelte";

  let exchanges = [];
  let banks = [];
  let fiatAmount = 0;

  let result = false;
  let rates = [];

  const submitForm = async () => {
    const payload = {
      exchange: exchanges,
      fiat: "RUB",
      bank: banks,
      tradeType: "BUY",
      fiatAmount: fiatAmount,
    };

    try {
      const response = await fetch(
        "http://localhost:5000/api/GetUsdtFiatRate",
        {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(payload),
        }
      );

      const data = await response.json();
      result = data.result;
      rates = data.rates;
    } catch (error) {
      console.error("Error submitting form:", error);
    }
  };

  const resetForm = () => {
    exchanges = [];
    banks = [];
    fiatAmount = 0;
    result = false;
    rates = [];
  };

  $: exchangeString = exchanges.join(", ");
  $: bankString = banks.join(", ");

  onMount(() => {
    console.log("Component mounted");
  });
</script>

<style>
  /* Add your styles here if needed */
</style>

<main>
  <form on:submit|preventDefault="{submitForm}" on:reset="{resetForm}">
    <fieldset>
      <legend>Exchanges</legend>
      <label>
        <input type="checkbox" bind:checked="{exchanges}" value="Binance" />
        Binance
      </label>
      <label>
        <input type="checkbox" bind:checked="{exchanges}" value="Bybit" /> ByBit
      </label>
    </fieldset>

    <fieldset>
      <legend>Banks</legend>
      <label>
        <input type="checkbox" bind:checked="{banks}" value="Sber" /> Sber
      </label>
      <label>
        <input type="checkbox" bind:checked="{banks}" value="Tinkoff" /> Tinkoff
      </label>
    </fieldset>

    <fieldset>
      <legend>Parameters</legend>
      <label>
        Fiat Amount: <input type="number" bind:value="{fiatAmount}" />
      </label>
    </fieldset>

    <fieldset>
      <legend>Actions</legend>
      <button type="submit">Submit</button>
      <button type="reset">Reset</button>
    </fieldset>
  </form>

  {#if result}
  <table>
    <thead>
      <tr>
        <th>Exchange</th>
        <th>Bank</th>
        <th>Price</th>
      </tr>
    </thead>
    <tbody>
      {#each rates as { exchange, bank, price }}
      <tr>
        <td>{exchange}</td>
        <td>{bank}</td>
        <td>{price}</td>
      </tr>
      {/each}
    </tbody>
  </table>
  {/if}
</main>
```

Для запуска:

- Установить Node.js
- `npx dev-server`

## Этап 3. Добавление SQLite базы данных

- Цель: Хранение данных о производимых обменах

Задача. Нужно создать таблицу аналогичную той что сейчас используется в Google Sheets

## Этап 3.1. Хранение данных об операциях перевода

### Метод для добавления новой записи операции перевода

Пример запроса

```request
POST api/exchange
{
    "exchange": "Binance"
    "fromBank": "Sber"
    "fromFiat": "RUB",
    "fromAmount": 1234,
    "toBank": "BKB"
    "toFiat": "THB",
    "toAmount": 1234,
    "client": "John Doe",
    "comissionUsdt": 0.0,
    "state": "new"
}
```

Ответ возвращает созданный объект + идентификатор. Пример:

```json
{
  "result": true,
  "exchange": {
    "id": 1,
    "exchange": "Binance"
    "fromBank": "Sber"
    "fromFiat": "RUB",
    "fromAmount": 1234,
    "toBank": "BKB"
    "toFiat": "THB",
    "toAmount": 1234,
    "client": "John Doe",
    "comissionUsdt": 0.0,
    "state": "new"
  }
}
```

### Метод для получения списка операции перевода

```request
GET api/exchange
```

```json
{
  "result": true,
  "exchanges": [
    {
      "id": 1,
      "exchange": "Binance"
      "fromBank": "Sber"
      "fromFiat": "RUB",
      "fromAmount": 1234,
      "toBank": "BKB"
      "toFiat": "THB",
      "toAmount": 1234,
      "client": "John Doe",
      "comissionUsdt": 0.0,
      "state": "new"
    }
  ]
}
```

### Метод для получения существующей операции перевода

```request
GET api/exchange/1
```

В ответе возвращается объект с идентификатором 1. Пример:

```json
{
  "result": true,
  "exchage": {
    "id": 1,
    "exchange": "Binance"
    "fromBank": "Sber"
    "fromFiat": "RUB",
    "fromAmount": 1234,
    "toBank": "BKB"
    "toFiat": "THB",
    "toAmount": 1234,
    "client": "John Doe",
    "comissionUsdt": 0.0,
    "state": "new"
  }
}
```

### Метод для изменения существующей операции перевода

Пример запроса:

```request
PUT /api/exchange/1
{
  "exchange": "Binance"
  "fromBank": "Sber"
  "fromFiat": "RUB",
  "fromAmount": 1234,
  "toBank": "BKB"
  "toFiat": "THB",
  "toAmount": 1234,
  "client": "John Doe",
  "comissionUsdt": 0.5,
  "state": "complete"
}
```

В ответе возващается измененный объект
