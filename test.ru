setVar["i";0;"int"]
while[i<=10]
    setVar["j";0;"int"]
    while[j<=10]
        if[(j*i)%2==0]
            print["j*i is even: {j*i}"]
        else[]
            print["j*i is odd: {j*i}"]
        end[]
        setVar["j";j+1]
    end[]
    setVar["i";i+1]
end[]