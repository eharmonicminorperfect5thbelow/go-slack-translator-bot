function doGet(e) {
  const text = e.parameter.text
  const from = e.parameter.from
  const to = e.parameter.to
  
  const translated = LanguageApp.translate(text, from, to)
  
  return ContentService.createTextOutput(translated)
}