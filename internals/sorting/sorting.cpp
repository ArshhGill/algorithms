#include <iostream>
#include "vector"
using namespace std;

// PADHAI ARC


vector<int> bubbleSort(vector<int> l){
    while (true){
        int c{0};
        for (int i = 0; i < l.size()-1; i ++){
            if (l[i] > l[i+1]){
                swap(l[i], l[i+1]);
                c++;
            }
        }
        if (c == 0){
            break;
        }
    }
    return l;
}

vector<int> insertionSort(vector<int> l){
    return l;
}

void bubbleSortRunner(vector<int> l){

    vector<int> bubbleSortReturn = bubbleSort(l);
    for (int i = 0; i < bubbleSortReturn.size(); i++){
        cout << bubbleSortReturn[i] << " ";
    }
}



int main(){
    vector<int> l = {1,5,3,7,3,6,8,43,0};

    for (int i = 0; i < l.size(); i++){
        cout << l[i] << " ";
    }
    cout << endl;

    bubbleSortRunner(l);

    return 0;
}

