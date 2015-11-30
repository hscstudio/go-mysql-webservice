# go-mysql-webservice
Another example go webservice with mysql data

## Prepare Database
You should have install mysql in Your server or PC tseting

Create database in mysql for example `restdb`, with structure:
```
CREATE TABLE IF NOT EXISTS `employee` (
  `id` int(11) NOT NULL,
  `name` varchar(50) NOT NULL,
  `email` varchar(100) NOT NULL,
  `birthday` date DEFAULT NULL,
  `photo` varchar(255) DEFAULT NULL,
  `created_at` int(11) DEFAULT NULL,
  `updated_at` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=8;
```

Add sampel data
```
INSERT INTO `employee` (`id`, `name`, `email`, `birthday`, `photo`, `created_at`, `updated_at`) VALUES
(1, 'Jayda Douglas', 'britchie@yahoo.com', '1983-12-18', 'images/3c6bb1a0f9559dc5f35b5c4ff0af62a1.jpg', 1448838415, NULL),
(2, 'Arne Cummings', 'bulah.christiansen@barrows.com', '1976-06-03', 'images/7aa5030f593d9fdbaf83928477581342.jpg', 1448838422, NULL),
(3, 'Ludwig Daniel', 'lorenz97@donnelly.net', '1937-07-25', 'images/16fe84f059901ac62575d0878233ff45.jpg', 1448838436, NULL),
(4, 'Devan Prosacco', 'qpagac@kiehn.com', '1969-11-14', 'images/113fa3eb20213b3a95f35f27e0ca5dc1.jpg', 1448838443, NULL),
(5, 'Jesse Sauer', 'emmitt26@gmail.com', '1983-12-18', 'images/978cb1271da894dbd3795e35de188585.jpg', 1448838458, NULL),
(6, 'Antonietta Adams Jr.', 'jordi.stark@yahoo.com', '1928-02-29', 'images/b102b35a0483c4a55f4f19b79834952a.jpg', 1448838465, NULL),
(7, 'Corine Grady', 'cleveland95@yahoo.com', '1983-12-18', 'images/c92658fcd87b5db8c88304c4ffc94ca5.jpg', 1448838466, NULL);
```

## Step by step
1. You should have install go
https://golang.org/doc/install

2. Via terminal, change Your current directory to src, then

3. Download Source
You may use git
```
git clone https://github.com/hscstudio/go-mysql-webservice
```
or You can download manually

4. Install dependency
```
cd go-mysql-webservice
go get
```
5. Run Go
```
go run rest.go
```
6. Access in browser 
http://localhost:8080/employee

Enjoy

Hafid Mukhlasin