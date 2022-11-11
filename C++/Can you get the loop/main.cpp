#include <algorithm>

using namespace std;

int getLoopSize(Node *startNode)
{
  vector<Node *> nodes;
  Node *currentNode = startNode;
  nodes.push_back(currentNode);
  while (currentNode->getNext() != nullptr)
  {
    currentNode = currentNode->getNext();
    auto loopStart = find(nodes.begin(), nodes.end(), currentNode);
    if (loopStart != nodes.end())
    {
      int index = loopStart - nodes.begin();
      return nodes.size() - index;
    }
    nodes.push_back(currentNode);
  }
  return 0;
}