# bcgo-w9

# Frequently Asked Questions (FAQ)

---

## ▶ Github
### Authentication

---

**✔ SSH**
> create new ssh key

```bash
# go to root folder
cd ~
# go to ssh folder
mkdir .ssh && cd .ssh

# save new ssh key
ssh-keygen -t rsa -b 4096 -C "your_email@example.com"
```

> config ssh config file
```bash
touch config
```
open the config file and add the following
```bash
Host *
  AddKeysToAgent yes
  UseKeychain yes
  IdentityFile ~/.ssh/id_rsa
```

---

**✔ Github**
> copy the SSH key
```bash
pbcopy < ~/.ssh/id_rsa.pub
```

> go to your Github account and add the SSH key

---

**✔ Test**

We test the SSH connection to Github (the identity file is specified in the config file)

> test the SSH connection
```bash
ssh -T git@github.com
```

In order to test other identities, you can use the following command
```bash
ssh -i ~/.ssh/key -T git@github.com
```

<br>

Up to now you have configured the SSH key (public / private), indicate ssh which identity file to use and added the SSH key (public) to your Github account.
We can use an optional service to manage the SSH keys (cache) and avoid to enter the phrase each time we use the SSH key.

---

**✔ SSH-Agent - Optional**

initialize ssh-agent as a service of ssh keys management (cache) (optional)

> start the ssh-agent in the background
```bash
eval "$(ssh-agent -s)"
```

> add your SSH private key to the ssh-agent
```bash
ssh-add -K ~/.ssh/id_rsa
```

<br>
<br>
<br>

## ▶ Terminal
in progress...

<br>
<br>
<br>

## ▶ Go
### Go Routines
Cuando trabajamos con concurrencia, una de las formas de comunicar los procesos es mediante canales . La otra forma es via "shared memory" aunque aplica otros mecanismos y es mas riesgoso. Un canal implica la comunicación entre 2 entidades, un 'Sender' y un 'Receiver'. Si un proceso termina de realizar su tarea, puede dejar sus resultados o algún mensaje en el canal, que le será de utilidad a la entidad o proceso que lo reciba. Las reglas que aplican dependerá de las características en las que funcione el canal.

Existen 2 tipos: Unbuffered Channel y Buffered Channel
> Unbuffered Channel: es un canal que no tiene un tamaño definido, de comunicacion directa. Esto implica que el 'Sender' y el 'Receiver' deben estar sincronizados para que la comunicación sea exitosa.
> - Caso 1: Si el 'Sender' envía un mensaje y el 'Receiver' no lo recibe, el 'Sender' se quedará esperando hasta que el 'Receiver' lo reciba y el proceso se estancará, pues el sender no podra continuar con otras tareas.
> - Caso 2: Si el 'Receiver' recibe un mensaje y el 'Sender' no lo envía, el 'Receiver' se quedará esperando hasta que el 'Sender' lo envíe. Tambien genera un estancamiento, pues el 'Receiver' no podrá continuar con otras tareas.
>
> El estancamiento se conoce como "Deadlock"

![alt text](/assets/goroutines-unbuffered.png "Unbuffered Channel")

> Buffered Channel: es un canal que tiene un tamaño definido, de comunicacion indirecta. Esto implica que el 'Sender' y el 'Receiver' no necesariamente deben estar sincronizados para que la comunicación sea exitosa.
> - Caso 1: Si el 'Sender' envía un mensaje y el 'Receiver' no lo recibe, el 'Sender' no se quedará esperando hasta que el 'Receiver' lo reciba, pues el canal tiene un tamaño definido y puede almacenar mensajes. El 'Sender' podrá continuar con otras tareas.
> - Caso 2: Si el 'Receiver' recibe un mensaje y el 'Sender' no lo envía, el 'Receiver' no se quedará esperando hasta que el 'Sender' lo envíe, pues el canal tiene un tamaño definido y puede almacenar mensajes. El 'Receiver' podrá continuar con otras tareas hasta que el canal quede vacío.

![alt text](/assets/goroutines-buffered.png "Buffered Channel")
