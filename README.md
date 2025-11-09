#Описание проекта.
Проект создан как выпускная работа на курсе Hive, с целью закрепления полученных знаний и объединение их всех в структурировонном приложении. 
Проект представляет из себя приложение TODO листа с возможностями создания, редактирования, удаления задач, аутенфикация по паролю. 


# http://localhost:7540/

# settings.go
package tests

 var Port = 7540
 
 var DBFile = "../scheduler.db"

 var FullNextDate = false

 var Search = false

 var Token = `secret123`

# Для сборки образа
sudo docker build -t todo-scheduler .

# Для запуска без пароля
sudo docker run -p 7540:7540 todo-scheduler

# Запуск с паролем
sudo docker run -p 7540:7540 -e TODO_PASSWORD="secret123" todo-scheduler
