const fs = require("fs");

const input = fs.readFileSync("./input.txt").toString().split("\n");

const numbers = [
  { k: "one", v: 1 },
  { k: "two", v: "2" },
  { k: "three", v: "3" },
  { k: "four", v: "4" },
  { k: "five", v: "5" },
  { k: "six", v: "6" },
  { k: "seven", v: "7" },
  { k: "eight", v: "8" },
  { k: "nine", v: "9" },
];

const regexify = (x) => {
  return parseInt(x.replace(/^(?:^.*?(\d).*(\d).*|.*(\d).*$)/, "$1$2$3$3"), 10);
};

const main = () => {
  const numbersButWithLengths = numbers.map((x) => ({ ...x, l: x.k.length }));

  let total = 0;

  for (const line of input) {
    if (!line.length) continue;

    let finalLine = "";

    for (let position = 0; position < line.length; position++) {
      const numberToAdd = parseInt(line[position], 10);

      if (!Number.isNaN(numberToAdd)) {
        finalLine += numberToAdd.toString();
        continue;
      }

      for (const nb of numbersButWithLengths) {
        if (line[position] === nb.k[0]) {
          let shouldAddNumber = true;
          for (let nbPosition = 1; nbPosition < nb.l; nbPosition++) {
            if (nb.k[nbPosition] !== line[position + nbPosition]) {
              shouldAddNumber = false;
              break;
            }
          }

          if (shouldAddNumber) {
            const numberToAdd = nb.v;
            finalLine += numberToAdd;
            break;
          }
        }
      }
    }

    const calibrationValues = regexify(finalLine);

    total += calibrationValues;
  }

  console.log(total);
};

main();
