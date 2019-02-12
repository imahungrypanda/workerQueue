for (let i = 0; i < 50; i++) {
  const body = new FormData();
  body.append("delay", "1s");
  body.append("name", 1);
  body.append("type", 'test');

  fetch('http://127.0.0.1:8000/work', {
    method: 'POST',
    mode: 'no-cors',
    headers: {
      "Content-Type": "form-data",
    },
    body
  }).then(response => console.log(response))
}
