class Deque():
	def __init__(self):
		self.items = []
	
	def isEmpty(self):
		return self.items == []

	def addFront(self,item):
		self.items.insert(0,item)

	def removeFront(self):
		self.items.pop(0)

	def addRear(self,item):
		self.items.append(item)

	def removeRear(self):
		self.items.pop()

	def size(self):
		return len(self.items)

def main():
	d = Deque()
	print(d.isEmpty())

	d.addFront(0)
	d.addFront(1)
	d.addRear(2)
	d.addRear(3)

	print(d.items)

	d.removeFront()
	print(d.items)
	d.removeRear()
	print(d.items)
main()
