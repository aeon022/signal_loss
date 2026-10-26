// Loads the built page in a real browser, starts a run in fast mode in each language
// and plays the first choice. Fails on any page error or if the game never answers.
// Usage: npm run smoke   (needs `astro preview` running; CHROME_BIN picks the browser)
import puppeteer from "puppeteer-core";

const base = process.env.SMOKE_URL ?? "http://localhost:4321/signal_loss/";
const executablePath = process.env.CHROME_BIN ?? "google-chrome";
const browser = await puppeteer.launch({ executablePath, headless: true, args: ["--no-sandbox", "--lang=en-US"] });

const screen = (page) => page.evaluate(() => document.querySelector(".xterm-rows")?.innerText ?? "");
async function until(page, text, what) {
  try {
    await page.waitForFunction((t) => document.querySelector(".xterm-rows")?.innerText.includes(t), { timeout: 60000 }, text);
  } catch {
    throw new Error(`${what}: never saw "${text}". Screen:\n${await screen(page)}`);
  }
}

try {
  for (const [path, lang, choose, next] of [["", "en", "choose", "Chapter 2"], ["de/", "de", "wählen", "Kapitel 2"]]) {
    const page = await browser.newPage();
    await page.evaluateOnNewDocument(() => { localStorage.setItem("sl_lang", "en"); localStorage.setItem("sl_cookie_ack", "1"); }); // no language redirect, no notice over the buttons
    const errors = [];
    page.on("pageerror", (e) => errors.push(e.message));
    await page.goto(base + path, { waitUntil: "networkidle0" });
    if (!(await page.$(`#launch-panel [data-value="${lang}"][aria-pressed="true"]`))) throw new Error(`${lang}: wrong language preselected`);
    await page.click('#launch-panel [data-group="speed"] [data-value="fast"]');
    await page.click("#launch-btn");
    await until(page, choose, `${lang} first screen`);
    await page.keyboard.press("1");
    await until(page, "▶", `${lang} outcome`); // outcome waits for a key
    await page.keyboard.press("Space");
    await until(page, next, `${lang} chapter 2`);
    if (errors.length) throw new Error(`${lang}: page errors: ${errors.join("; ")}`);
    console.log(`ok ${lang}`);
    await page.close();
  }
  // A fatal first choice must end the run, count it and offer the share text.
  const page = await browser.newPage();
  await page.evaluateOnNewDocument(() => { localStorage.clear(); localStorage.setItem("sl_lang", "en"); localStorage.setItem("sl_cookie_ack", "1"); });
  await page.goto(base, { waitUntil: "networkidle0" });
  await page.click('#launch-panel [data-group="speed"] [data-value="fast"]');
  await page.click("#launch-btn");
  await until(page, "choose", "death run");
  await page.keyboard.press("3");
  await until(page, "continue", "death outcome");
  await page.keyboard.press("Space");
  await until(page, "GAME OVER", "game over");
  await page.keyboard.press("Space");
  await page.waitForSelector("#share:not(.hidden)", { timeout: 20000 });
  const shared = await page.$eval("#share-text", (e) => e.textContent);
  if (!shared.includes("death #1")) throw new Error(`share text: ${shared}`);
  console.log("ok death + share");
} finally {
  await browser.close();
}
