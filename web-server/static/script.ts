// script.ts
const calculateBtn = document.getElementById('calculateBtn') as HTMLButtonElement;
const result = document.getElementById('result') as HTMLParagraphElement;

calculateBtn.addEventListener('click', () => {
  const num1 = parseFloat((document.getElementById('num1') as HTMLInputElement).value);
  const num2 = parseFloat((document.getElementById('num2') as HTMLInputElement).value);

  if (isNaN(num1) || isNaN(num2)) {
    result.textContent = "Please enter valid numbers!";
  } else {
    const sum = num1 + num2;
    result.textContent = `The sum is: ${sum}`;
  }
});
