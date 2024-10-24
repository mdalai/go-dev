// script.ts
var calculateBtn = document.getElementById('calculateBtn');
var result = document.getElementById('result');
calculateBtn.addEventListener('click', function () {
    var num1 = parseFloat(document.getElementById('num1').value);
    var num2 = parseFloat(document.getElementById('num2').value);
    if (isNaN(num1) || isNaN(num2)) {
        result.textContent = "Please enter valid numbers!";
    }
    else {
        var sum = num1 + num2;
        result.textContent = "The sum is: ".concat(sum);
    }
});
