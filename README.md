# iBox-Server

---

## storage.bin 格式

### [Folder]

    - Id string
    - Mame string
    - Share uint8
        [0000 0000]
              ____ Public Permission
        ____ ShareUser Permission
    - ShareUser []string
    - SubFolders []Folder

---

## Permission 權限

- ### NO_PERMISSION

  - 無法檢視

- ### READ_ONLY

  - 可以檢視檔案、資料夾

- ### ALL

  - 可以編輯、新增、修改、刪除(不含權限設定)
