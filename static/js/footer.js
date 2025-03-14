// @ts-nocheck
ScrollReveal().reveal(".footer", {
  reset: true,
  delay: 200,
});
// <p id="distortion-text" class="text font-bold">Have A Great Day!</p>

var elem = document.getElementById("distortion-text");

var inlineText = elem.innerText;

var text = new Blotter.Text(elem.innerText, {
  family: "'Wudoo Mono'",
  size: 22,
  weight: 600,
  paddingLeft: 120,
  paddingRight: 120,
});

var material = new Blotter.ChannelSplitMaterial();

material.uniforms.uApplyBlur.value = true;
material.uniforms.uOffset.value = 0.125;

var blotter = new Blotter(material, {
  texts: text,
});

var scope = blotter.forText(text);

elem.innerText = "";

scope.appendTo(elem);
