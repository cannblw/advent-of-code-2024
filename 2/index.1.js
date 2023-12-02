const fs = require("fs");
const readline = require("readline");

const main = async () => {
  const limits = {
    red: 12,
    green: 13,
    blue: 14,
  };

  const fileStream = fs.createReadStream("input.txt");

  const rl = readline.createInterface({
    input: fileStream,
    crlfDelay: Infinity,
  });

  let total = 0;

  for await (const line of rl) {
    let isGamePossible = true;

    const gameAndGameDescription = line.split(": ");
    const gameId = gameAndGameDescription[0].split(" ")[1];
    const gameDescriptions = gameAndGameDescription[1].split("; ");

    game_loop: for (const gameDescription of gameDescriptions) {
      const plays = gameDescription.split(", ").map((x) => x.split(" "));

      for (const [numberStr, color] of plays) {
        if (parseInt(numberStr, 10) > limits[color]) {
          isGamePossible = false;
          break game_loop;
        }
      }
    }

    if (isGamePossible) {
      total += parseInt(gameId, 10);
    }
  }

  console.log(total);
};

main();
