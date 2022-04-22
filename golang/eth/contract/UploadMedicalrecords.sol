pragma solidity ^0.8.0;

import "https://github.com/OpenZeppelin/openzeppelin-contracts/blob/master/contracts/token/ERC20/IERC20.sol";
import "https://github.com/OpenZeppelin/openzeppelin-contracts/blob/master/contracts/token/ERC20/ERC20.sol";
//import "./a.sol"

contract UploadMedicalrecords is ERC20{

   address public A11Smile;

  constructor() ERC20("AS","AS") { 

     A11Smile = msg.sender;

    _mint(msg.sender,20000 * 10 ** uint(18));

  }



   event Addtokens( address indexed owner, uint indexed amount);//增加代币事件

   event Uploadmedicaldata(address indexed user,string indexed route,address indexed soliciter ); //用户上传医疗信息事件

   event gainerUploadmedicaldata(string indexed route,address indexed soliciter ); //征求者上传医疗信息事件

   event reward(address indexed owner,string indexed route,uint indexed price); //征求者奖励事件

   event A11GiveETH(address indexed user1, address indexed user2,uint indexed price);//A11Smile给新用户发ETH事件

   event EthgetAs(address indexed user1, address indexed user2,uint indexed price);//ETH购买AS事件





//增加代币

function mint1( uint amount) public {

   _mint(msg.sender,amount ** uint(18));

   emit Addtokens(msg.sender,amount);

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



//征求者发布医疗数据信息

struct gainer_Medicalinformation1{

  string[] PictureRoute;//病历名称

  address soliciter;

  mapping(string => gainer_Medicalinformation2 ) gainerAmedicalInformation2;

}

struct gainer_Medicalinformation2{

  uint min;  //设置最小奖励金额

  uint max;  //设置最大奖励金额

  string Medicalrecordrequirements;//病历要求

  string Requirementdescription;//需求描述

}



//给合约转ETH

function toContract() public payable{



}



//用户上传体检报告

struct MedicalExaminationReportstrucr{

  mapping(string => string) file;

}



//用户上传病历信息

struct Medicalinformation{

   mapping(string => string) file;

}





//用户上传医疗数据

struct user_MedicalCertificate{

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

mapping (address => user_MedicalCertificate) private userAmedicalInformation;



//mapping用户对应上传体检报告

mapping(address => MedicalExaminationReportstrucr) user_MedicalExaminationReportstrucrURL;


mapping(address => string[]) public user_MedicalExaminationReportstrucrName;




//mapping用户对应上传病历信息

mapping(address => Medicalinformation) user_MedicalinformationURL;

mapping(address => string[]) public user_MedicalinformationrName;





//mapping征求者对应发布医疗数据信息

mapping (address => gainer_Medicalinformation1) private gainerAmedicalInformation1;



//mapping征求者对应查询医疗数据信息

mapping(address => user_MedicalCertificate) private gainAmedicalInformation;



//mapping征求者审核自己的病历信息

mapping(address => gainer_examine ) private gainerExamine;

mapping(string => user_MedicalCertificate ) private gainerExamineString;



//mapping查询医疗数据是否通过

mapping(string => gainer_examine ) private userExamineString;



//初始化由A11Smile给用户ETH

function A11SmileGiveETH(address payable Startuse) public payable  {

  require(msg.sender == A11Smile, "you not can transfer");

  require(msg.sender.balance > msg.value,"Your account doesn't have so much ETH");

  Startuse.transfer(msg.value);

  emit A11GiveETH(msg.sender,Startuse,msg.value);

}





//设置用户

function user_Adduser(address _user) public{

  require(!user_IsUser(_user),"User is exit");

    Ausers[_user].user= _user;

    Ausers[_user].exit= true;

}



//设置征求者 

function gainer_setDoctor(address soliciter) public{

  require(!gainer_isDoctor(soliciter),"Soliciter is exit");

  Adoctor[soliciter].soliciter=soliciter;

  Adoctor[soliciter].exist=true;

}





//用户上传证书

function user_AddMedicalInformation(string memory Proute,address _soliciter) public user_upTime(msg.sender) {

      require(user_IsUser(msg.sender),"User no exist");

      require(gainer_isDoctor(_soliciter),"Soliciter no exist");

      userAmedicalInformation[msg.sender].user=msg.sender;

      userAmedicalInformation[msg.sender].PictureRoute.push(Proute);

      userAmedicalInformation[_soliciter].soliciter=_soliciter;

      // emit Uploadmedicaldata(msg.sender,Proute,_soliciter);

}



//用户上传体检报告

function user_UPMedicalExaminationReport(string memory name,string memory url) public {

  user_MedicalExaminationReportstrucrURL[msg.sender].file[name] = url;

  user_MedicalExaminationReportstrucrName[msg.sender].push(name);

}

// 查看用户的体检报告量

function viewMedicalExaminationReportCount() public view returns(uint){
    return user_MedicalExaminationReportstrucrName[msg.sender].length;
}




//用户上传病历信息

function user_UPMedicalinformation(string memory name,string memory url) public {

    user_MedicalinformationURL[msg.sender].file[name] = url;

    user_MedicalinformationrName[msg.sender].push(name);

}

// 查看用户的病历信息量

function viewMedicalinformationCount() public view returns(uint){
    return user_MedicalinformationrName[msg.sender].length;
}



//征求者发布自己所需的医疗信息 

function gainer_AddMedicalInformation(uint min,uint max,string memory MedicalName,string memory _Medicalrecordrequirements,string memory _Requirementdescription) public {

  require(gainer_isDoctor(msg.sender),"Soliciter no exist");

      gainerAmedicalInformation1[msg.sender].PictureRoute.push(MedicalName);

      gainerAmedicalInformation1[msg.sender].gainerAmedicalInformation2[MedicalName].min= min;

      gainerAmedicalInformation1[msg.sender].gainerAmedicalInformation2[MedicalName].max= max;

      gainerAmedicalInformation1[msg.sender].gainerAmedicalInformation2[MedicalName].Medicalrecordrequirements= _Medicalrecordrequirements;

      gainerAmedicalInformation1[msg.sender].gainerAmedicalInformation2[MedicalName].Requirementdescription= _Requirementdescription;

      gainerAmedicalInformation1[msg.sender].soliciter=msg.sender;

      emit gainerUploadmedicaldata(MedicalName,msg.sender);

}



//征求者查看自己发布的医疗信息

function gainer_SeeOwnMedicalInformation(string memory MedicalName) public view returns(string[] memory,uint,uint,string memory,string memory,address){

    return (gainerAmedicalInformation1[msg.sender].PictureRoute,

    gainerAmedicalInformation1[msg.sender].gainerAmedicalInformation2[MedicalName].min,

    gainerAmedicalInformation1[msg.sender].gainerAmedicalInformation2[MedicalName].max,

    gainerAmedicalInformation1[msg.sender].gainerAmedicalInformation2[MedicalName].Medicalrecordrequirements,

    gainerAmedicalInformation1[msg.sender].gainerAmedicalInformation2[MedicalName].Requirementdescription,

    gainerAmedicalInformation1[msg.sender].soliciter);

}



//查询用户是否存在

function user_IsUser(address _user) public view returns(bool){

  return Ausers[_user].exit;

}



//查询征求者是否存在

function gainer_isDoctor(address soliciter)public view returns(bool){

  return Adoctor[soliciter].exist;

}



//用户查看用户上传的证书

function user_ViewMedicalRecords() public view returns(string[] memory,address soliciter){

  require(user_IsUser(msg.sender),"patient no exist");

  return (userAmedicalInformation[msg.sender].PictureRoute,userAmedicalInformation[msg.sender].soliciter);

}



//用户查看用户上传的体检报告

function user_ViewMedicalExaminationReport(string memory name) public view returns(string memory){

  return user_MedicalExaminationReportstrucrURL[msg.sender].file[name];

}





//用户查看用户上传的医疗信息

function user_ViewMedicalInformation(string memory name) public view returns(string memory){

  return  user_MedicalinformationURL[msg.sender].file[name];

}





//征求者查看用户上传给自己的证书

function  gainer_doctorsee(address soliciter) public view returns(address,string[] memory,address){

  return (gainAmedicalInformation[soliciter].user,gainAmedicalInformation[soliciter].PictureRoute,gainAmedicalInformation[soliciter].soliciter);

}



//设置代币地址

function A11Smile_setErc(address ercaddress_) public{

     ercaddress=ercaddress_;

}



//征求者审核,奖励

function gainer_whether(address _soliciter,string memory PictureRoute,bool whether,uint ercnum_) public gainer_examineTime(_soliciter){

        gainerExamine[_soliciter].soliciter=_soliciter;

        gainerExamine[_soliciter].user=gainerExamineString[PictureRoute].user;

        gainerExamine[_soliciter].pRoute=PictureRoute;

        require(_soliciter==gainerExamineString[PictureRoute].soliciter,"The medical data does not belong to the soliciter");

       if(whether==true){

         gainerExamine[_soliciter].whether = true;

        gainer_transfer_AS(gainerExamineString[PictureRoute].soliciter,gainerExamineString[PictureRoute].user,ercnum_);

       }else{

         revert("Failed to pass the review");

       }

       emit reward(_soliciter,PictureRoute,ercnum_);

}





//征求者奖励,奖励AS代币

function gainer_transfer_AS(address from1 ,address to1,uint quantity1) public {

  ERC20.transferFrom(from1,to1,quantity1);

}



//查看证书是否通过审核

function seeMedicaldata(string memory PictureRoute) public view returns(string memory,bool){

    return(PictureRoute,userExamineString[PictureRoute].whether);

} 



//查看当前账户的ETH

function getUserETH() public view returns(uint) {

  return msg.sender.balance;

}



//查看当前账户的AS

function getUserAS() public view returns(uint) {

  return balanceOf(msg.sender);

}


}
