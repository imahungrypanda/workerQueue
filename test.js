function request(obj) {
  return new Promise(((resolve, reject) => {
    let xhr = new XMLHttpRequest();

    xhr.open(
      obj.method || 'POST',
      obj.url,
      true,
    );

    xhr.onload = () => {
      if (xhr.status >= 200 && xhr.status < 300) {
        resolve(xhr.response);
      } else {
        reject(xhr.statusText);
      }
    };

    xhr.onerror = () => {
      reject(xhr.statusText);
    };

    xhr.send(obj.body);
  }));
}


setTimeout(() => {
  for (let i = 0; i < 50; i++) {
    const body = new FormData();
    body.append("delay", "1s");
    body.append("name", i);

    request({ url: "http://127.0.0.1:8000/work", body }).then(response => console.log("response"))
    // request({ url: "https://dev.rigminder.com/tcli/", body }).then(response => console.log("response"))
  }

}, 10)
