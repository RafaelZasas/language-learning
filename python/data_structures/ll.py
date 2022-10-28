#!/usr/local/bin/python3
class Node:
	def __init__(self, value=None):
		self.value = value
		self.next = None

class LinkedList:
	def __init__(self):
		self.head = None

def add(l: LinkedList):
	
