const fs = require("fs");
const readline = require("readline");

const main = async () => {
  const fileStream = fs.createReadStream("input.txt");

  const rl = readline.createInterface({
    input: fileStream,
    crlfDelay: Infinity,
  });

  let total = 0;

  for await (const line of rl) {
    const gameAndGameDescription = line.split(": ");
    const gameDescriptions = gameAndGameDescription[1].split("; ");

    const maxItemsInPlay = {};

    for (const gameDescription of gameDescriptions) {
      const plays = gameDescription.split(", ").map((x) => x.split(" "));

      for (const [numberStr, color] of plays) {
        const numberOfCubes = parseInt(numberStr, 10);

        maxItemsInPlay[color] = Math.max(
          maxItemsInPlay[color] ?? 0,
          numberOfCubes
        );
      }
    }

    const power = Object.values(maxItemsInPlay).reduce(
      (prev, current) => prev * current,
      1
    );

    total += power;
  }

  console.log(total);
};

main();
