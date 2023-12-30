package main

import "github.com/wallybarnum/bigskyutil/cmd/bigskyutil"

func main() {
    bigskyutil.Execute()
}

/*
json_obj: {"cmd":"read","path":"/presets/pst0.json"}
sysex msg to send:  SysExType data: 00 01 55 18 01 50 00 7B 22 63 6D 64 22 3A 22 72 65 61 64 22 2C 22 70 61 74 68 22 3A 22 2F 70 72 65 73 65 74 73 2F 70 73 74 30 2E 6A 73 6F 6E 22 7D
waiting for response
RX msg
RX sysex
   0  00 01 55 18 01 50 45                              ..U..PE         
sysex msg to send:  SysExType data: 00 01 55 18 01 50 01 01 00
waiting for response
RX msg
RX sysex


   0  00 01 55 18 01 50 45 00  7B 22 69 6E 66 6F 22 00  ..U..PE.{"info".
  16  20 3A 20 7B 22 70 72 00  6F 64 75 63 74 22 3A 00   : {"pr.oduct":.
  32  22 42 69 67 53 6B 79 00  20 4D 58 22 2C 22 6E 00  "BigSky. MX","n.
  48  61 6D 65 22 3A 22 45 00  56 45 52 59 44 41 59 00  ame":"E.VERYDAY.
  64  22 2C 22 69 73 64 75 00  61 6C 22 20 3A 20 30 00  ","isdu.al" : 0.
  80  2C 22 66 6F 72 6D 61 00  74 22 20 3A 20 31 20 00  ,"forma.t" : 1 .
  96  7D 2C 22 73 74 61 74 00  65 22 20 3A 20 7B 22 00  },"stat.e" : {".
 112  49 4E 46 20 4C 41 54 00  43 48 22 20 3A 20 30 00  INF LAT.CH" : 0.
 128  2C 22 42 4F 4F 53 54 00  22 20 3A 20 38 2C 22 00  ,"BOOST." : 8,".
 144  50 45 52 53 49 53 54 00  22 20                    PERSIST."       

   0  00 01 55 18 01 50 45 00  3A 20 30 2C 22 45 58 00  ..U..PE.: 0,"EX.
  16  50 20 53 45 54 55 50 00  22 20 3A 20 30 2C 22 00  P SETUP." : 0,".
  32  44 55 41 4C 22 20 3A 00  20 30 2C 22 63 68 61 00  DUAL" :. 0,"cha.
  48  6E 6E 65 6C 22 20 3A 00  20 5B 7B 22 65 78 70 00  nnel" :. [{"exp.
  64  72 65 73 73 69 6F 6E 00  22 20 3A 20 5B 7B 22 00  ression." : [{".
  80  6B 6E 6F 62 22 3A 22 00  44 65 63 61 79 22 2C 00  knob":".Decay",.
  96  22 68 65 65 6C 22 20 00  3A 20 30 2C 22 74 6F 00  "heel" .: 0,"to.
 112  65 22 20 3A 20 30 20 00  7D 2C 7B 22 6B 6E 6F 00  e" : 0 .},{"kno.
 128  62 22 3A 22 50 72 65 00  2D 44 65 6C 61 79 22 00  b":"Pre.-Delay".
 144  2C 22 68 65 65 6C 22 00  20 3A                    ,"heel". :      
 
   0  00 01 55 18 01 50 45 00  20 30 2C 22 74 6F 65 00  ..U..PE. 0,"toe.
  16  22 20 3A 20 30 20 7D 00  2C 7B 22 6B 6E 6F 62 00  " : 0 }.,{"knob.
  32  22 3A 22 4D 69 78 22 00  2C 22 68 65 65 6C 22 00  ":"Mix".,"heel".
  48  20 3A 20 30 2C 22 74 00  6F 65 22 20 3A 20 30 00   : 0,"t.oe" : 0.
  64  20 7D 2C 7B 22 6B 6E 00  6F 62 22 3A 22 54 6F 00   },{"kn.ob":"To.
  80  6E 65 22 2C 22 68 65 00  65 6C 22 20 3A 20 30 00  ne","he.el" : 0.
  96  2C 22 74 6F 65 22 20 00  3A 20 30 20 7D 2C 7B 00  ,"toe" .: 0 },{.
 112  22 6B 6E 6F 62 22 3A 00  22 4D 6F 64 22 2C 22 00  "knob":."Mod",".
 128  68 65 65 6C 22 20 3A 00  20 30 2C 22 74 6F 65 00  heel" :. 0,"toe.
 144  22 20 3A 20 30 20 7D 00  20 5D                    " : 0 }. ]      
   0  00 01 55 18 01 50 45 00  2C 22 54 59 50 45 22 00  ..U..PE.,"TYPE".
  16  20 3A 20 39 2C 22 53 00  54 41 54 45 22 20 3A 00   : 9,"S.TATE" :.
  32  20 31 2C 22 50 61 72 00  61 6D 31 22 3A 22 4C 00   1,"Par.am1":"L.
  48  4F 57 20 45 4E 44 22 00  2C 22 50 61 72 61 6D 00  OW END".,"Param.
  64  32 22 3A 22 43 4F 4C 00  4F 52 22 2C 22 44 65 00  2":"COL.OR","De.
  80  63 61 79 22 20 3A 20 00  31 30 32 30 2C 22 50 00  cay" : .1020,"P.
  96  72 65 2D 44 65 6C 61 00  79 22 20 3A 20 30 2C 00  re-Dela.y" : 0,.
 112  22 4D 69 78 22 20 3A 00  20 31 34 37 2C 22 54 00  "Mix" :. 147,"T.
 128  6F 6E 65 22 20 3A 20 00  38 34 2C 22 4D 6F 64 00  one" : .84,"Mod.
 144  22 20 3A 20 31 38 39 00  2C 22                    " : 189.,"      
   0  00 01 55 18 01 50 45 00  4C 4F 57 20 45 4E 44 00  ..U..PE.LOW END.
  16  22 20 3A 20 38 2C 22 00  43 4F 4C 4F 52 22 20 00  " : 8,".COLOR" .
  32  3A 20 31 2C 22 4F 55 00  54 50 55 54 20 4C 45 00  : 1,"OU.TPUT LE.
  48  56 45 4C 22 20 3A 20 00  31 36 2C 22 50 41 4E 00  VEL" : .16,"PAN.
  64  22 20 3A 20 38 2C 22 00  49 4E 46 20 4D 4F 44 00  " : 8,".INF MOD.
  80  45 22 20 3A 20 30 20 00  7D 2C 7B 22 65 78 70 00  E" : 0 .},{"exp.
  96  72 65 73 73 69 6F 6E 00  22 20 3A 20 5B 7B 22 00  ression." : [{".
 112  6B 6E 6F 62 22 3A 22 00  44 65 63 61 79 22 2C 00  knob":".Decay",.
 128  22 68 65 65 6C 22 20 00  3A 20 30 2C 22 74 6F 00  "heel" .: 0,"to.
 144  65 22 20 3A 20 30 20 00  7D 2C                    e" : 0 .},      
   0  00 01 55 18 01 50 45 00  7B 22 6B 6E 6F 62 22 00  ..U..PE.{"knob".
  16  3A 22 50 72 65 2D 44 00  65 6C 61 79 22 2C 22 00  :"Pre-D.elay",".
  32  68 65 65 6C 22 20 3A 00  20 30 2C 22 74 6F 65 00  heel" :. 0,"toe.
  48  22 20 3A 20 30 20 7D 00  2C 7B 22 6B 6E 6F 62 00  " : 0 }.,{"knob.
  64  22 3A 22 4D 69 78 22 00  2C 22 68 65 65 6C 22 00  ":"Mix".,"heel".
  80  20 3A 20 30 2C 22 74 00  6F 65 22 20 3A 20 30 00   : 0,"t.oe" : 0.
  96  20 7D 2C 7B 22 6B 6E 00  6F 62 22 3A 22 54 6F 00   },{"kn.ob":"To.
 112  6E 65 22 2C 22 68 65 00  65 6C 22 20 3A 20 30 00  ne","he.el" : 0.
 128  2C 22 74 6F 65 22 20 00  3A 20 30 20 7D 2C 7B 00  ,"toe" .: 0 },{.
 144  22 6B 6E 6F 62 22 3A 00  22 4D                    "knob":."M      
   0  00 01 55 18 01 50 45 00  6F 64 22 2C 22 68 65 00  ..U..PE.od","he.
  16  65 6C 22 20 3A 20 30 00  2C 22 74 6F 65 22 20 00  el" : 0.,"toe" .
  32  3A 20 30 20 7D 20 5D 00  2C 22 54 59 50 45 22 00  : 0 } ].,"TYPE".
  48  20 3A 20 31 2C 22 53 00  54 41 54 45 22 20 3A 00   : 1,"S.TATE" :.
  64  20 30 2C 22 50 61 72 00  61 6D 31 22 3A 22 44 00   0,"Par.am1":"D.
  80  49 46 46 55 53 49 4F 00  4E 22 2C 22 50 61 72 00  IFFUSIO.N","Par.
  96  61 6D 32 22 3A 22 4C 00  4F 57 20 45 4E 44 22 00  am2":"L.OW END".
 112  2C 22 44 65 63 61 79 00  22 20 3A 20 31 30 30 00  ,"Decay." : 100.
 128  30 2C 22 50 72 65 2D 00  44 65 6C 61 79 22 20 00  0,"Pre-.Delay" .
 144  3A 20 31 32 37 2C 22 00  4D 69                    : 127,".Mi      
   0  00 01 55 18 01 50 45 00  78 22 20 3A 20 31 32 00  ..U..PE.x" : 12.
  16  37 2C 22 54 6F 6E 65 00  22 20 3A 20 31 32 37 00  7,"Tone." : 127.
  32  2C 22 4D 6F 64 22 20 00  3A 20 31 32 37 2C 22 00  ,"Mod" .: 127,".
  48  44 49 46 46 55 53 49 00  4F 4E 22 20 3A 20 38 00  DIFFUSI.ON" : 8.
  64  2C 22 4C 4F 57 20 45 00  4E 44 22 20 3A 20 30 00  ,"LOW E.ND" : 0.
  80  2C 22 45 4E 53 45 4D 00  42 4C 45 22 20 3A 20 00  ,"ENSEM.BLE" : .
  96  30 2C 22 4F 55 54 50 00  55 54 20 4C 45 56 45 00  0,"OUTP.UT LEVE.
 112  4C 22 20 3A 20 31 36 00  2C 22 50 41 4E 22 20 00  L" : 16.,"PAN" .
 128  3A 20 38 2C 22 49 4E 00  46 20 4D 4F 44 45 22 00  : 8,"IN.F MODE".
 144  20 3A 20 30 20 7D 20 00  5D 20                     : 0 } .]       

   0  00 01 55 18 01 50 45 00  7D 20 7D                 ..U..PE.} }     

   0  00 01 55 18 01 50 46                              ..U..PF         

   0  00 01 55 18 01 50 46                              ..U..PF         

   0  00 01 55 18 01 50 46                              ..U..PF      










{"info" : {"product":"BigSky MX","name":"EVERYDAY","isdual" : 0,"format" : 1 },"state" : {"INF LATCH" : 0,"BOOST" : 8,"PERSIST": 0,"EXP SETUP" : 0,"DUAL" : 0,"channel" : [{"expression" : [{"knob":"Decay","heel" : 0,"toe" : 0 },{"knob":"Pre-Delay","heel" : 0,"toe" : 0 },{"knob":"Mix","heel" : 0,"toe" : 0 },{"knob":"Tone","heel" : 0,"toe" : 0 },{"knob":"Mod","heel" : 0,"toe" : 0 } ],"TYPE" : 9,"STATE" : 1,"Param1":"LOW END","Param2":"COLOR","Decay" : 1020,"Pre-Delay" : 0,"Mix" : 147,"Tone" : 84,"Mod" : 189,"LOW END" : 8,"COLOR" : 1,"OUTPUT LEVEL" : 16,"PAN" : 8,"INF MODE" : 0 },{"expression" : [{"knob":"Decay","heel" : 0,"toe" : 0 },{"knob":"Pre-Delay","heel" : 0,"toe" : 0 },{"knob":"Mix","heel" : 0,"toe" : 0 },{"knob":"Tone","heel" : 0,"toe" : 0 },{"knob":"Mod","heel" : 0,"toe" : 0 } ],"TYPE" : 1,"STATE" : 0,"Param1":"DIFFUSION","Param2":"LOW END","Decay" : 1000,"Pre-Delay" : 127,"Mix" : 127,"Tone" : 127,"Mod" : 127,"DIFFUSION" : 8,"LOW END" : 0,"ENSEMBLE" : 0,"OUTPUT LEVEL" : 16,"PAN" : 8,"INF MODE" : 0 } ] } }    


{
    "info": {
        "product": "BigSky MX",
        "name": "EVERYDAY",
        "isdual": 0,
        "format": 1
    },
    "state": {
        "INF LATCH": 0,
        "BOOST": 8,
        "PERSIST": 0,
        "EXP SETUP": 0,
        "DUAL": 0,
        "channel": [
            {
                "expression": [
                    {
                        "knob": "Decay",
                        "heel": 0,
                        "toe": 0
                    },
                    {
                        "knob": "Pre-Delay",
                        "heel": 0,
                        "toe": 0
                    },
                    {
                        "knob": "Mix",
                        "heel": 0,
                        "toe": 0
                    },
                    {
                        "knob": "Tone",
                        "heel": 0,
                        "toe": 0
                    },
                    {
                        "knob": "Mod",
                        "heel": 0,
                        "toe": 0
                    }
                ],
                "TYPE": 9,
                "STATE": 1,
                "Param1": "LOW END",
                "Param2": "COLOR",
                "Decay": 1020,
                "Pre-Delay": 0,
                "Mix": 147,
                "Tone": 84,
                "Mod": 189,
                "LOW END": 8,
                "COLOR": 1,
                "OUTPUT LEVEL": 16,
                "PAN": 8,
                "INF MODE": 0
            },
            {
                "expression": [
                    {
                        "knob": "Decay",
                        "heel": 0,
                        "toe": 0
                    },
                    {
                        "knob": "Pre-Delay",
                        "heel": 0,
                        "toe": 0
                    },
                    {
                        "knob": "Mix",
                        "heel": 0,
                        "toe": 0
                    },
                    {
                        "knob": "Tone",
                        "heel": 0,
                        "toe": 0
                    },
                    {
                        "knob": "Mod",
                        "heel": 0,
                        "toe": 0
                    }
                ],
                "TYPE": 1,
                "STATE": 0,
                "Param1": "DIFFUSION",
                "Param2": "LOW END",
                "Decay": 1000,
                "Pre-Delay": 127,
                "Mix": 127,
                "Tone": 127,
                "Mod": 127,
                "DIFFUSION": 8,
                "LOW END": 0,
                "ENSEMBLE": 0,
                "OUTPUT LEVEL": 16,
                "PAN": 8,
                "INF MODE": 0
            }
        ]
    }
}



{"info" : {"product":"BigSky MX","name":"EVERYDAY","isdual" : 0,"format" : 1 },"state" : {"INF LATCH" : 0,"BOOST" : 8,"PERSIST" : 0,"EXP SETUP" : 0,"DUAL" : 0,"channel" : [{"expression" : [{"knob":"Decay","heel" : 0,"toe" : 0 },{"knob":"Pre-Delay","heel" : 0,"toe" : 0 },{"knob":"Mix","heel" : 0,"toe" : 0 },{"knob":"Tone","heel" : 0,"toe" : 0 },{"knob":"Mod","heel" : 0,"toe" : 0 } ],"TYPE" : 9,"STATE" : 1,"Param1":"LOW END","Param2":"COLOR","Decay" : 1020,"Pre-Delay" : 0,"Mix" : 147,"Tone" : 84,"Mod" : 189,"LOW END" : 8,"COLOR" : 1,"OUTPUT LEVEL" : 16,"PAN" : 8,"INF MODE" : 0 },{"expression" : [{"knob":"Decay","heel" : 0,"toe" : 0 },{"knob":"Pre-Delay","heel" : 0,"toe" : 0 },{"knob":"Mix","heel" : 0,"toe" : 0 },{"knob":"Tone","heel" : 0,"toe" : 0 },{"knob":"Mod","heel" : 0,"toe" : 0 } ],"TYPE" : 1,"STATE" : 0,"Param1":"DIFFUSION","Param2":"LOW END","Decay" : 1000,"Pre-Delay" : 127,"Mix" : 127,"Tone" : 127,"Mod" : 127,"DIFFUSION" : 8,"LOW END" : 0,"ENSEMBLE" : 0,"OUTPUT LEVEL" : 16,"PAN" : 8,"INF MODE" : 0 } ] } }
*/