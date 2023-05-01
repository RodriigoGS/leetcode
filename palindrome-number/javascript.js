const isPalindrome = function(x) {
    let inverse = 0;
    let copy = x;

    while (copy > 0) {
        inverse = inverse * 10 + (copy % 10);
        copy = ~~(copy / 10);
    }

    return inverse === x;
};