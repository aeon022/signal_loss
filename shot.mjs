import puppeteer from "puppeteer-core";
const b = await puppeteer.launch({ executablePath: "/Applications/Google Chrome.app/Contents/MacOS/Google Chrome", headless: true, args: ["--lang=en-US"] });
for (const [name, vp] of [["desk", {width:1440,height:900}], ["phone", {width:390,height:844,isMobile:true,hasTouch:true}]]) {
  const p = await b.newPage(); await p.setViewport(vp);
  p.on("pageerror", e => console.log("PAGEERR", e.message));
  await p.evaluateOnNewDocument(() => { localStorage.clear(); localStorage.setItem("sl_lang","en"); localStorage.setItem("sl_cookie_ack","1"); });
  await p.goto("http://localhost:4322/signal_loss/", { waitUntil: "networkidle0" });
  await p.click('#launch-panel [data-group="lang"] [data-value="en"]');
  await p.click("#launch-btn");
  await p.waitForFunction(() => document.querySelector(".xterm-rows")?.innerText.includes("uplink"), {timeout:30000});
  await new Promise(r=>setTimeout(r,800));
  console.log(name, "locked:", await p.evaluate(() => document.documentElement.classList.contains("play-lock")));
  await p.screenshot({ path: `${name}.png` });
  if (name==="desk") { await p.evaluate(()=>window.scrollTo(0,99999)); console.log("scrollY after wheel-ish:", await p.evaluate(()=>scrollY)); }
}
await b.close();
