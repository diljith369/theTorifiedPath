from stem.control import Controller

print('Connecting to tor')

with Controller.from_port() as controller:
  controller.authenticate(password="SecretPass")
  response = controller.create_ephemeral_hidden_service({80: 5000}, await_publication = True)
  print("Our onion domain is %s.onion:80, pointing to localhost:5000" % response.service_id)
  try:
    input("Press Enter, or ^C to exit")
  finally:
    print("Shutting down our hidden service")