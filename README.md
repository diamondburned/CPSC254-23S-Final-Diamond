# CPSC254-23S-Final-Diamond

9. See below.

   1. Repository URL: https://github.com/diamondburned/CPSC254-23S-Final-Diamond

   2. Screenshots:

      Generating a new public/private key pair:
  
      ![](https://0x0.st/HNgT.png)
  
      Print the public key and copy it:
  
      ![](https://0x0.st/HNgA.png)
  
      Adding the newly generated public key to GitHub:
  
      ![](https://0x0.st/HNgm.png)
  
      Key is added:
  
      ![](https://0x0.st/HNga.png)
  
      Cloning using SSH:
  
      ![](https://0x0.st/HNgB.png)

10. ```bash
    mkdir -p final/{src,data,config}
    mv main.cpp header.h final/src
    mv mainDB.db backupDB.db final/data
    mv config.yml final/config
    
    tree final # print the directory structure
    ```

    Before:

    ![](https://0x0.st/HNgM.png)

    After:

    ![](https://0x0.st/HNgu.png)

11. For question 11, I decided to implement this in Go and vanilla JS. Reasons
    being:

    - I like writing Go, so that's what I chose.
    - I want to start with the bare minimum frontend, so I wrote HTML pages and
      added a little bit of JS to make it work.

    The backend has no database. There are polymorphic interfaces describing the
    required stores, and I used a mock store for the demo, which is just some
    simple Go code and a hashmap. This is the quickest way to get a working
    storage without having to set up a database.

    The backend follows basic Go conventions, so it's pretty easy to read. I
    also split the code into multiple packages:

    - `frontend/` contains the frontend HTML code.
    - `backend/` contains the backend server and API routes. It consumes the
      `bank` package
    - `bank/` contains the bank models and stores. It describes the skeleton
      (interfaces) of the bank; everything else either implements the interfaces
      or uses them.
    - `main.go` is the entry point of the program.
    - `mockbank.go` contains the mock implementation of the bank store.

    The frontend is just a few HTML pages with some JS code to make it talk to
    the backend. I used [MVP.css](https://andybrewer.github.io/mvp/) to make it
    look a lot nicer.
