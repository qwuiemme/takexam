function validate() {
  form = document.getElementById("form");
  login = document.getElementById("login").value;
  pass = document.getElementById("password").value;
  passconf = document.getElementById("password-confirm").value;

  if (login.length < 4) {
    changeErrorDivStatus(false, "Длина логина меньше 4 символов!");
    return;
  }
  if (pass.length < 6) {
    changeErrorDivStatus(false, "Длина пароля меньше 6 символов!");
    return;
  }
  if (pass !== passconf) {
    changeErrorDivStatus(false, "Пароли не совпадают!");
    return;
  }

  form.submit();
}

function changeErrorDivStatus(hidden, text) {
  let errordiv = document.getElementById("error");
  let errortext = document.getElementById("error-text");

  errordiv.setAttribute("aria-hidden", String(hidden));
  errortext.innerText = text;
}
