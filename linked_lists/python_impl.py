class Nodes:
	def __init__(self,data):
		self.data = data
		self.next = None

def main():
	linkedList = Nodes(1)
	n2 = Nodes(2)
	n3 = Nodes(3)
	n2.next = n3
	linkedList.next = n2

	s = linkedList
	while True:
		print(s)
		if s.next == None:
			print("End of linked list reached")
			break
		s = s.next

main()
