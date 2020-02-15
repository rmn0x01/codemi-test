#variabel global
loker = []
lowest_index = 0

#inisialisasi n loker
def init_loker():
    global loker
    while(len(loker)==0):
        command = input().split()
        try:
            if(command[0].lower() == 'init'):
                if(int(command[1])<1):
                    print('inisialisasi loker harus lebih dari 0')
                else:
                    loker = [['kosong',''] for i in range(int(command[1]))]
                    print('Berhasil membuat loker dengan jumlah', command[1])
            else:
                print('Jumlah loker belum ditentukan')
        except:
            print('inisialisasi loker: init [angka]')
        

def get_status():
    global loker
    print("No Loker\tTipe Identitas\tNo Identitas")
    for i in range(len(loker)):
        print(str(i+1)+'\t'+'\t'+loker[i][0]+'\t\t'+loker[i][1])

def get_lowest_index():
    global lowest_index
    updated = False
    for i in range(lowest_index,len(loker)):
        if(loker[i][1]==''):
            updated = True
            lowest_index = i
            break
    if(not updated):
        lowest_index = -1
    return lowest_index
    

def input_data(tipe,nomor):
    global loker, lowest_index
    if(lowest_index>=0):
        loker[lowest_index][0] = tipe
        loker[lowest_index][1] = nomor
        print("Kartu identitas tersimpan di loker nomor",str(lowest_index+1))
        get_lowest_index()
    else:
        print("Maaf loker sudah penuh")

def delete_data(no_loker):
    global loker, lowest_index
    if(no_loker.isdigit() and no_loker!='0'):
        no_loker = int(no_loker)
        if(no_loker>len(loker)):
            print("Tidak ada loker dengan nomor", no_loker)
        elif(loker[no_loker-1][1]==''):
            print("Loker sudah kosong")
        else:
            loker[no_loker-1][0]='kosong'
            loker[no_loker-1][1]=''
            if ((lowest_index==-1) or ((no_loker-1)<lowest_index)):
                lowest_index=(no_loker-1)
            print("Loker nomor",no_loker,"berhasil dikosongkan")
    else:
        print("Nomor loker harus berupa angka positif")

def find_data(no_identitas):
    global loker
    found = []
    for i in range(len(loker)):
        if(loker[i][1] == no_identitas):
            found.append(str(i+1))
    if(len(found)==0):
        print("Nomor identitas tidak ditemukan")
    else:
        print("Kartu identitas tersebut berada di loker nomor "+", ".join(found))

def main():
    init_loker()
    #navigasi utama
    running = True
    while(running):
        command = input().split()
        if(len(command)<1):
            print("Available commands: status, input, leave, find, exit")
            continue
        if(command[0].lower() == 'status'):
            if(len(command)>1):
                print('usage: status')
            else:
                get_status()
        elif(command[0].lower() == 'input'):
            if(len(command)==3):
                tipe = command[1]
                nomor = command[2]
                input_data(tipe,nomor)
            else:
                print('usage: input [tipe identitas] [nomor identitas]')
        elif(command[0].lower() == 'leave'):
            if(len(command) == 2):
                no_loker = command[1]
                delete_data(no_loker)
            else:
                print('usage: leave [nomor loker]')
        elif(command[0].lower() == 'find'):
            if(len(command) == 2):
                no_identitas = command[1]
                find_data(no_identitas)
            else:
                print("usage: find [no. identitas]")
        elif(command[0].lower() == 'init'):
            print("Jumlah loker sudah diinisialisasi")
        elif(command[0].lower() == 'exit'):
            print("Closing...")
            running = False
        else:
            print("Available commands: status, input, leave, find, exit")




if __name__== "__main__":
    main()