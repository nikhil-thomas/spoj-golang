// Title             : 'Feynman'
// SPOJ ID           : 'SAMER08F'
// Description       : 'https://www.spoj.com/problems/SAMER08F/'
// Author            : 'Nikhil Thomas'
// Created On        : 'August 11, 2018'
// Last Modified On  : 'August 11, 2018'
//----------::::::::::----------::::::::::----------::::::::::----------//

// SPOJ Specifications ---------::::::::::----------::::::::::----------//
// Added by     : Diego Satoba
// Date         : 2008-11-23
// Time limit   : 0.165s
// Source limit : 50000B
// Memory limit : 1536MB
// Cluster      : Cube (Intel G860)
// Languages    : ASM64 C CPP C++ 4.3.2 FORTRAN JAVA PAS-GPC PAS-FPC
// Resource     : South American Regional Contests 2008
//----------::::::::::----------::::::::::----------::::::::::----------//

#include <iostream>

using namespace std;

struct node {
  int val;
  node *next;
};

// computeSquareCount computes nunmber of suqares in this problem
// num of squares = num of 1 sided squares + num of 2 sided squares + ... + num of n sided squares
// num of squares = n^2 + (n-1)^2 + ... + 1^2
int computeSquareCount(int n) {
	int squares = 0;
	for (int i = n; i >= 0; i--) {
		squares += i * i;
	}
	return squares;
};

int main() {
  node *top = NULL;
  node *bottom = NULL;

  while (true) {
    int val = 0;
    cin >> val;
    if (val == 0) {
      break;
    }
    node *item = new(node);
    item->val = val;
    item->next = NULL;
    if (top == NULL) {
      top = item;
    } else {
      bottom->next = item;
    }
    bottom = item;

  }

  node *item = top;

  while (item != NULL) {
    cout << computeSquareCount(item->val) << "\n";
    node *temp = item;
    item = item->next;
    delete(temp);
  }
}
