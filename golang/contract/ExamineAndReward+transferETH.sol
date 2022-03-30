pragma solidity ^0.8.0;



import "https://github.com/OpenZeppelin/openzeppelin-contracts/blob/master/contracts/token/ERC20/IERC20.sol";  

import "https://github.com/OpenZeppelin/openzeppelin-contracts/blob/master/contracts/token/ERC20/ERC20.sol"; 

//import "./a.sol"



contract UploadMedicalrecords is ERC20{

address public A11Smile;

  constructor() ERC20("AS","AS") { 

​    A11Smile = msg.sender;

​    _mint(msg.sender,20000 * 10 ** uint(decimals()));

  }





//用户 

struct user_Users {

  address user;

  bool exit;

}



//征求者地址

struct Soliciter{

  address soliciter;

  bool exist;

}



//代币地址

address ercaddress = address(this);



//上传医疗数据信息时间

uint public lastUpdateTime;



//审核医疗数据时间

uint public examineTime;



//用户上传医疗数据信息

struct user_Medicalinformation{

  address user;

  string[] PictureRoute;

  address soliciter;

}



//征求者审核

struct gainer_examine{

  address soliciter;

  address user;

  string pRoute;

  bool  whether;

}



//modified上传时间

modifier user_upTime(address _user){

  lastUpdateTime = block.timestamp;



  _;

}



//modified审核时间

modifier gainer_examineTime(address soliciter){

  examineTime = block.timestamp;



  _;

}



//mapping对应用户

mapping (address => user_Users)  private Ausers;



//mapping对应征求者

mapping(address => Soliciter) private Adoctor;  



//mapping用户对应上传医疗数据信息

mapping (address => user_Medicalinformation) private userAmedicalInformation;



//mapping征求者对应查询医疗数据信息

mapping(address => user_Medicalinformation) private gainAmedicalInformation;



//mapping征求者审核自己的病历信息

mapping(address => gainer_examine ) private gainerExamine;

mapping(string => user_Medicalinformation ) private gainerExamineString;



//mapping查询医疗数据是否通过

mapping(string => gainer_examine ) private userExamineString;



//初始化由A11Smile给用户ETH

function adcc(address payable Startuse) public payable  {

  require(msg.sender == A11Smile, "you not can transfer");

  require(msg.sender.balance > msg.value,"Your account doesn't have so much ETH");

  Startuse.transfer(msg.value);

}



//设置用户

function user_Adduser(address _user) public{

  require(!user_IsUser(_user),"User is exit");

​    Ausers[_user].user= _user;

​    Ausers[_user].exit= true;

}



//设置征求者 

function gainer_setDoctor(address soliciter) public{

  require(!gainer_isDoctor(soliciter),"Soliciter is exit");

  Adoctor[soliciter].soliciter=soliciter;

  Adoctor[soliciter].exist=true;

}





//用户上传医疗信息

function user_AddMedicalInformation(string memory Proute,address _soliciter) public user_upTime(msg.sender) {

​      require(user_IsUser(msg.sender),"User no exist");

​      require(gainer_isDoctor(_soliciter),"Soliciter no exist");

​      userAmedicalInformation[msg.sender].user=msg.sender;

​      userAmedicalInformation[msg.sender].PictureRoute.push(Proute);

​      userAmedicalInformation[_soliciter].soliciter=_soliciter;

}

  

//查询用户是否存在

function user_IsUser(address _user) public view returns(bool){

  return Ausers[_user].exit;

}



//查询征求者是否存在

function gainer_isDoctor(address soliciter)public view returns(bool){

  return Adoctor[soliciter].exist;

}



//用户查看用户上传的病历

function user_ViewMedicalRecords(address _user) public view returns(string[] memory,address soliciter){

  require(user_IsUser(_user),"patient no exist");

  return (userAmedicalInformation[_user].PictureRoute,userAmedicalInformation[_user].soliciter);

}



//征求者查看用户上传给自己的医疗信息

function  gainer_doctorsee(address soliciter) public view returns(address,string[] memory,address){

  return (gainAmedicalInformation[soliciter].user,gainAmedicalInformation[soliciter].PictureRoute,gainAmedicalInformation[soliciter].soliciter);

}



//设置代币地址

function setErc(address ercaddress_) public{

​     ercaddress=ercaddress_;

}



//征求者审核,奖励

function gainer_whether(address _soliciter,string memory PictureRoute,bool whether,uint ercnum_) public gainer_examineTime(_soliciter){

​        gainerExamine[_soliciter].soliciter=_soliciter;

​        gainerExamine[_soliciter].user=gainerExamineString[PictureRoute].user;

​        gainerExamine[_soliciter].pRoute=PictureRoute;

​        require(_soliciter==gainerExamineString[PictureRoute].soliciter,"The medical data does not belong to the soliciter");

​       if(whether==true){

​         gainerExamine[_soliciter].whether = true;

​        _transfer_AS(gainerExamineString[PictureRoute].soliciter,gainerExamineString[PictureRoute].user,ercnum_);

​       }else{

​         revert("Failed to pass the review");

​       }



}





//征求者奖励,奖励AS代币

function _transfer_AS(address from1 ,address to1,uint quantity1) public {

  ERC20.transferFrom(from1,to1,quantity1);

}



//查看医疗数据是否通过审核

function  seeMedicaldata(string memory PictureRoute) public view returns(string memory,bool){

​    return(PictureRoute,userExamineString[PictureRoute].whether);

} 



//查看当前账户的ETH

function getUserBalance(address user) public view returns(uint) {

  return user.balance;

}



}