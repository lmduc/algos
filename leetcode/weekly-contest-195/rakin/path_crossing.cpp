#include <iostream>
#include <unordered_map>
#include <string>

using namespace std;

int rowAdd[] = {1, -1, 0, 0};
int colAdd[] = {0, 0, 1, -1};

class Solution {
public:

    int pickInd(char c) {
        if(c == 'N') {
            return 0;
        }
        if(c == 'E') {
            return 2;
        }
        if(c == 'S') {
            return 1;
        }
        return 3;
    }

    string toStr(int r, int c) {
        return to_string(r) + " " + to_string(c);
    }

    bool isPathCrossing(string path) {
        unordered_map<string, int> map;
        int rr = 0;
        int cc = 0;
        map[toStr(rr, cc)]++;

        for(char c : path) {
            int ind = pickInd(c);
            rr += rowAdd[ind];
            cc += colAdd[ind];
            
            string str = toStr(rr, cc);
            map[str]++;
            if(map[str] > 1) return true;
        }

        return false;   
    }
};

int main() {
    cout << "test" << endl;
}