<?xml version="1.0" encoding="utf-8"?>
<!DOCTYPE eagle SYSTEM "eagle.dtd">
<eagle version="7.6.0">
<drawing>
<settings>
<setting alwaysvectorfont="no"/>
<setting verticaltext="up"/>
</settings>
<grid distance="0.1" unitdist="inch" unit="inch" style="lines" multiple="1" display="no" altdistance="0.01" altunitdist="inch" altunit="inch"/>
<layers>
<layer number="1" name="Top" color="4" fill="1" visible="no" active="no"/>
<layer number="2" name="Route2" color="1" fill="3" visible="no" active="no"/>
<layer number="3" name="Route3" color="4" fill="3" visible="no" active="no"/>
<layer number="4" name="Route4" color="1" fill="4" visible="no" active="no"/>
<layer number="5" name="Route5" color="4" fill="4" visible="no" active="no"/>
<layer number="6" name="Route6" color="1" fill="8" visible="no" active="no"/>
<layer number="7" name="Route7" color="4" fill="8" visible="no" active="no"/>
<layer number="8" name="Route8" color="1" fill="2" visible="no" active="no"/>
<layer number="9" name="Route9" color="4" fill="2" visible="no" active="no"/>
<layer number="10" name="Route10" color="1" fill="7" visible="no" active="no"/>
<layer number="11" name="Route11" color="4" fill="7" visible="no" active="no"/>
<layer number="12" name="Route12" color="1" fill="5" visible="no" active="no"/>
<layer number="13" name="Route13" color="4" fill="5" visible="no" active="no"/>
<layer number="14" name="Route14" color="1" fill="6" visible="no" active="no"/>
<layer number="15" name="Route15" color="4" fill="6" visible="no" active="no"/>
<layer number="16" name="Bottom" color="1" fill="1" visible="no" active="no"/>
<layer number="17" name="Pads" color="2" fill="1" visible="no" active="no"/>
<layer number="18" name="Vias" color="2" fill="1" visible="no" active="no"/>
<layer number="19" name="Unrouted" color="6" fill="1" visible="no" active="no"/>
<layer number="20" name="Dimension" color="15" fill="1" visible="no" active="no"/>
<layer number="21" name="tPlace" color="7" fill="1" visible="no" active="no"/>
<layer number="22" name="bPlace" color="7" fill="1" visible="no" active="no"/>
<layer number="23" name="tOrigins" color="15" fill="1" visible="no" active="no"/>
<layer number="24" name="bOrigins" color="15" fill="1" visible="no" active="no"/>
<layer number="25" name="tNames" color="7" fill="1" visible="no" active="no"/>
<layer number="26" name="bNames" color="7" fill="1" visible="no" active="no"/>
<layer number="27" name="tValues" color="7" fill="1" visible="no" active="no"/>
<layer number="28" name="bValues" color="7" fill="1" visible="no" active="no"/>
<layer number="29" name="tStop" color="7" fill="3" visible="no" active="no"/>
<layer number="30" name="bStop" color="7" fill="6" visible="no" active="no"/>
<layer number="31" name="tCream" color="7" fill="4" visible="no" active="no"/>
<layer number="32" name="bCream" color="7" fill="5" visible="no" active="no"/>
<layer number="33" name="tFinish" color="6" fill="3" visible="no" active="no"/>
<layer number="34" name="bFinish" color="6" fill="6" visible="no" active="no"/>
<layer number="35" name="tGlue" color="7" fill="4" visible="no" active="no"/>
<layer number="36" name="bGlue" color="7" fill="5" visible="no" active="no"/>
<layer number="37" name="tTest" color="7" fill="1" visible="no" active="no"/>
<layer number="38" name="bTest" color="7" fill="1" visible="no" active="no"/>
<layer number="39" name="tKeepout" color="4" fill="11" visible="no" active="no"/>
<layer number="40" name="bKeepout" color="1" fill="11" visible="no" active="no"/>
<layer number="41" name="tRestrict" color="4" fill="10" visible="no" active="no"/>
<layer number="42" name="bRestrict" color="1" fill="10" visible="no" active="no"/>
<layer number="43" name="vRestrict" color="2" fill="10" visible="no" active="no"/>
<layer number="44" name="Drills" color="7" fill="1" visible="no" active="no"/>
<layer number="45" name="Holes" color="7" fill="1" visible="no" active="no"/>
<layer number="46" name="Milling" color="3" fill="1" visible="no" active="no"/>
<layer number="47" name="Measures" color="7" fill="1" visible="no" active="no"/>
<layer number="48" name="Document" color="7" fill="1" visible="no" active="no"/>
<layer number="49" name="Reference" color="7" fill="1" visible="no" active="no"/>
<layer number="51" name="tDocu" color="7" fill="1" visible="no" active="no"/>
<layer number="52" name="bDocu" color="7" fill="1" visible="no" active="no"/>
<layer number="90" name="Modules" color="5" fill="1" visible="yes" active="yes"/>
<layer number="91" name="Nets" color="2" fill="1" visible="yes" active="yes"/>
<layer number="92" name="Busses" color="1" fill="1" visible="yes" active="yes"/>
<layer number="93" name="Pins" color="2" fill="1" visible="no" active="yes"/>
<layer number="94" name="Symbols" color="4" fill="1" visible="yes" active="yes"/>
<layer number="95" name="Names" color="7" fill="1" visible="yes" active="yes"/>
<layer number="96" name="Values" color="7" fill="1" visible="yes" active="yes"/>
<layer number="97" name="Info" color="7" fill="1" visible="yes" active="yes"/>
<layer number="98" name="Guide" color="6" fill="1" visible="yes" active="yes"/>
</layers>
<schematic xreflabel="%F%N/%S.%C%R" xrefpart="/%S.%C%R">
<libraries>
<library name="ETHER_BRKOUT">
<packages>
<package name="ETHER_BRKOUT">
<pad name="A-" x="-2" y="-1" drill="0.8"/>
<pad name="B-" x="0" y="-1" drill="0.8"/>
<pad name="C-" x="2" y="-1" drill="0.8"/>
<pad name="D-" x="4" y="-1" drill="0.8"/>
<pad name="A+" x="-3" y="1" drill="0.8"/>
<pad name="B+" x="-1" y="1" drill="0.8"/>
<pad name="C+" x="1" y="1" drill="0.8"/>
<pad name="D+" x="3" y="1" drill="0.8"/>
<text x="-1.5" y="2" size="0.8128" layer="21">B+</text>
<text x="0.5" y="2" size="0.8128" layer="21">C+</text>
<text x="2.5" y="2" size="0.8128" layer="21">D+</text>
<text x="-2.5" y="-3" size="0.8128" layer="21">A-</text>
<text x="-0.5" y="-3" size="0.8128" layer="21">B-</text>
<text x="1.5" y="-3" size="0.8128" layer="21">C-</text>
<text x="3.5" y="-3" size="0.8128" layer="21">D-</text>
<text x="-3.5" y="2" size="0.8128" layer="21">A+</text>
</package>
</packages>
<symbols>
<symbol name="ETHER_BRKOUT">
<pin name="A+" x="-7.62" y="-10.16" length="middle" rot="R90"/>
<pin name="A-" x="-5.08" y="-10.16" length="middle" rot="R90"/>
<pin name="B+" x="-2.54" y="-10.16" length="middle" rot="R90"/>
<pin name="B-" x="0" y="-10.16" length="middle" rot="R90"/>
<pin name="C+" x="2.54" y="-10.16" length="middle" rot="R90"/>
<pin name="C-" x="5.08" y="-10.16" length="middle" rot="R90"/>
<pin name="D+" x="7.62" y="-10.16" length="middle" rot="R90"/>
<pin name="D-" x="10.16" y="-10.16" length="middle" rot="R90"/>
<wire x1="-10.16" y1="-5.08" x2="12.7" y2="-5.08" width="0.254" layer="94"/>
<wire x1="12.7" y1="-5.08" x2="12.7" y2="7.62" width="0.254" layer="94"/>
<wire x1="12.7" y1="7.62" x2="-10.16" y2="7.62" width="0.254" layer="94"/>
<wire x1="-10.16" y1="7.62" x2="-10.16" y2="-5.08" width="0.254" layer="94"/>
<text x="-2.54" y="5.08" size="1.778" layer="94">&gt;NAME</text>
</symbol>
</symbols>
<devicesets>
<deviceset name="ETHER_BRKOUT" prefix="X">
<gates>
<gate name="G$1" symbol="ETHER_BRKOUT" x="0" y="0"/>
</gates>
<devices>
<device name="" package="ETHER_BRKOUT">
<connects>
<connect gate="G$1" pin="A+" pad="A+"/>
<connect gate="G$1" pin="A-" pad="A-"/>
<connect gate="G$1" pin="B+" pad="B+"/>
<connect gate="G$1" pin="B-" pad="B-"/>
<connect gate="G$1" pin="C+" pad="C+"/>
<connect gate="G$1" pin="C-" pad="C-"/>
<connect gate="G$1" pin="D+" pad="D+"/>
<connect gate="G$1" pin="D-" pad="D-"/>
</connects>
<technologies>
<technology name=""/>
</technologies>
</device>
</devices>
</deviceset>
</devicesets>
</library>
</libraries>
<attributes>
</attributes>
<variantdefs>
</variantdefs>
<classes>
<class number="0" name="default" width="0" drill="0">
</class>
</classes>
<parts>
<part name="X1" library="ETHER_BRKOUT" deviceset="ETHER_BRKOUT" device=""/>
<part name="X6" library="ETHER_BRKOUT" deviceset="ETHER_BRKOUT" device=""/>
<part name="X2" library="ETHER_BRKOUT" deviceset="ETHER_BRKOUT" device=""/>
<part name="X3" library="ETHER_BRKOUT" deviceset="ETHER_BRKOUT" device=""/>
<part name="X4" library="ETHER_BRKOUT" deviceset="ETHER_BRKOUT" device=""/>
<part name="X5" library="ETHER_BRKOUT" deviceset="ETHER_BRKOUT" device=""/>
<part name="X7" library="ETHER_BRKOUT" deviceset="ETHER_BRKOUT" device=""/>
<part name="X8" library="ETHER_BRKOUT" deviceset="ETHER_BRKOUT" device=""/>
<part name="X9" library="ETHER_BRKOUT" deviceset="ETHER_BRKOUT" device=""/>
<part name="X10" library="ETHER_BRKOUT" deviceset="ETHER_BRKOUT" device=""/>
</parts>
<sheets>
<sheet>
<plain>
</plain>
<instances>
<instance part="X1" gate="G$1" x="38.1" y="45.72"/>
<instance part="X6" gate="G$1" x="38.1" y="20.32" rot="MR180"/>
<instance part="X2" gate="G$1" x="63.5" y="45.72"/>
<instance part="X3" gate="G$1" x="63.5" y="20.32" rot="MR180"/>
<instance part="X4" gate="G$1" x="88.9" y="45.72"/>
<instance part="X5" gate="G$1" x="88.9" y="20.32" rot="MR180"/>
<instance part="X7" gate="G$1" x="114.3" y="45.72"/>
<instance part="X8" gate="G$1" x="114.3" y="20.32" rot="MR180"/>
<instance part="X9" gate="G$1" x="139.7" y="45.72"/>
<instance part="X10" gate="G$1" x="139.7" y="20.32" rot="MR180"/>
</instances>
<busses>
</busses>
<nets>
<net name="N$1" class="0">
<segment>
<pinref part="X6" gate="G$1" pin="A+"/>
<pinref part="X1" gate="G$1" pin="A+"/>
<wire x1="30.48" y1="30.48" x2="30.48" y2="35.56" width="0.1524" layer="91"/>
</segment>
</net>
<net name="N$2" class="0">
<segment>
<pinref part="X6" gate="G$1" pin="A-"/>
<pinref part="X1" gate="G$1" pin="A-"/>
<wire x1="33.02" y1="30.48" x2="33.02" y2="35.56" width="0.1524" layer="91"/>
</segment>
</net>
<net name="N$3" class="0">
<segment>
<pinref part="X6" gate="G$1" pin="B+"/>
<pinref part="X1" gate="G$1" pin="B+"/>
<wire x1="35.56" y1="30.48" x2="35.56" y2="35.56" width="0.1524" layer="91"/>
</segment>
</net>
<net name="N$4" class="0">
<segment>
<pinref part="X6" gate="G$1" pin="B-"/>
<pinref part="X1" gate="G$1" pin="C-"/>
<wire x1="38.1" y1="30.48" x2="43.18" y2="35.56" width="0.1524" layer="91"/>
</segment>
</net>
<net name="N$5" class="0">
<segment>
<pinref part="X6" gate="G$1" pin="C+"/>
<pinref part="X1" gate="G$1" pin="B-"/>
<wire x1="40.64" y1="30.48" x2="40.64" y2="33.02" width="0.1524" layer="91"/>
<wire x1="40.64" y1="33.02" x2="38.1" y2="35.56" width="0.1524" layer="91"/>
</segment>
</net>
<net name="N$6" class="0">
<segment>
<pinref part="X6" gate="G$1" pin="C-"/>
<pinref part="X1" gate="G$1" pin="C+"/>
<wire x1="43.18" y1="30.48" x2="43.18" y2="33.02" width="0.1524" layer="91"/>
<wire x1="43.18" y1="33.02" x2="40.64" y2="35.56" width="0.1524" layer="91"/>
</segment>
</net>
<net name="N$7" class="0">
<segment>
<pinref part="X6" gate="G$1" pin="D+"/>
<pinref part="X1" gate="G$1" pin="D+"/>
<wire x1="45.72" y1="30.48" x2="45.72" y2="35.56" width="0.1524" layer="91"/>
</segment>
</net>
<net name="N$8" class="0">
<segment>
<pinref part="X6" gate="G$1" pin="D-"/>
<pinref part="X1" gate="G$1" pin="D-"/>
<wire x1="48.26" y1="30.48" x2="48.26" y2="35.56" width="0.1524" layer="91"/>
</segment>
</net>
<net name="N$9" class="0">
<segment>
<pinref part="X3" gate="G$1" pin="A+"/>
<pinref part="X2" gate="G$1" pin="A+"/>
<wire x1="55.88" y1="30.48" x2="55.88" y2="35.56" width="0.1524" layer="91"/>
</segment>
</net>
<net name="N$10" class="0">
<segment>
<pinref part="X3" gate="G$1" pin="A-"/>
<pinref part="X2" gate="G$1" pin="A-"/>
<wire x1="58.42" y1="30.48" x2="58.42" y2="35.56" width="0.1524" layer="91"/>
</segment>
</net>
<net name="N$11" class="0">
<segment>
<pinref part="X3" gate="G$1" pin="B+"/>
<pinref part="X2" gate="G$1" pin="B+"/>
<wire x1="60.96" y1="30.48" x2="60.96" y2="35.56" width="0.1524" layer="91"/>
</segment>
</net>
<net name="N$12" class="0">
<segment>
<pinref part="X3" gate="G$1" pin="B-"/>
<pinref part="X2" gate="G$1" pin="C-"/>
<wire x1="63.5" y1="30.48" x2="68.58" y2="35.56" width="0.1524" layer="91"/>
</segment>
</net>
<net name="N$13" class="0">
<segment>
<pinref part="X3" gate="G$1" pin="C+"/>
<pinref part="X2" gate="G$1" pin="B-"/>
<wire x1="66.04" y1="30.48" x2="66.04" y2="33.02" width="0.1524" layer="91"/>
<wire x1="66.04" y1="33.02" x2="63.5" y2="35.56" width="0.1524" layer="91"/>
</segment>
</net>
<net name="N$14" class="0">
<segment>
<pinref part="X3" gate="G$1" pin="C-"/>
<pinref part="X2" gate="G$1" pin="C+"/>
<wire x1="68.58" y1="30.48" x2="68.58" y2="33.02" width="0.1524" layer="91"/>
<wire x1="68.58" y1="33.02" x2="66.04" y2="35.56" width="0.1524" layer="91"/>
</segment>
</net>
<net name="N$15" class="0">
<segment>
<pinref part="X3" gate="G$1" pin="D+"/>
<pinref part="X2" gate="G$1" pin="D+"/>
<wire x1="71.12" y1="30.48" x2="71.12" y2="35.56" width="0.1524" layer="91"/>
</segment>
</net>
<net name="N$16" class="0">
<segment>
<pinref part="X3" gate="G$1" pin="D-"/>
<pinref part="X2" gate="G$1" pin="D-"/>
<wire x1="73.66" y1="30.48" x2="73.66" y2="35.56" width="0.1524" layer="91"/>
</segment>
</net>
<net name="N$17" class="0">
<segment>
<pinref part="X5" gate="G$1" pin="A+"/>
<pinref part="X4" gate="G$1" pin="A+"/>
<wire x1="81.28" y1="30.48" x2="81.28" y2="35.56" width="0.1524" layer="91"/>
</segment>
</net>
<net name="N$18" class="0">
<segment>
<pinref part="X5" gate="G$1" pin="A-"/>
<pinref part="X4" gate="G$1" pin="A-"/>
<wire x1="83.82" y1="30.48" x2="83.82" y2="35.56" width="0.1524" layer="91"/>
</segment>
</net>
<net name="N$19" class="0">
<segment>
<pinref part="X5" gate="G$1" pin="B+"/>
<pinref part="X4" gate="G$1" pin="B+"/>
<wire x1="86.36" y1="30.48" x2="86.36" y2="35.56" width="0.1524" layer="91"/>
</segment>
</net>
<net name="N$20" class="0">
<segment>
<pinref part="X5" gate="G$1" pin="B-"/>
<pinref part="X4" gate="G$1" pin="C-"/>
<wire x1="88.9" y1="30.48" x2="93.98" y2="35.56" width="0.1524" layer="91"/>
</segment>
</net>
<net name="N$21" class="0">
<segment>
<pinref part="X5" gate="G$1" pin="C+"/>
<pinref part="X4" gate="G$1" pin="B-"/>
<wire x1="91.44" y1="30.48" x2="91.44" y2="33.02" width="0.1524" layer="91"/>
<wire x1="91.44" y1="33.02" x2="88.9" y2="35.56" width="0.1524" layer="91"/>
</segment>
</net>
<net name="N$22" class="0">
<segment>
<pinref part="X5" gate="G$1" pin="C-"/>
<pinref part="X4" gate="G$1" pin="C+"/>
<wire x1="93.98" y1="30.48" x2="93.98" y2="33.02" width="0.1524" layer="91"/>
<wire x1="93.98" y1="33.02" x2="91.44" y2="35.56" width="0.1524" layer="91"/>
</segment>
</net>
<net name="N$23" class="0">
<segment>
<pinref part="X5" gate="G$1" pin="D+"/>
<pinref part="X4" gate="G$1" pin="D+"/>
<wire x1="96.52" y1="30.48" x2="96.52" y2="35.56" width="0.1524" layer="91"/>
</segment>
</net>
<net name="N$24" class="0">
<segment>
<pinref part="X5" gate="G$1" pin="D-"/>
<pinref part="X4" gate="G$1" pin="D-"/>
<wire x1="99.06" y1="30.48" x2="99.06" y2="35.56" width="0.1524" layer="91"/>
</segment>
</net>
<net name="N$25" class="0">
<segment>
<pinref part="X8" gate="G$1" pin="A+"/>
<pinref part="X7" gate="G$1" pin="A+"/>
<wire x1="106.68" y1="30.48" x2="106.68" y2="35.56" width="0.1524" layer="91"/>
</segment>
</net>
<net name="N$26" class="0">
<segment>
<pinref part="X8" gate="G$1" pin="A-"/>
<pinref part="X7" gate="G$1" pin="A-"/>
<wire x1="109.22" y1="30.48" x2="109.22" y2="35.56" width="0.1524" layer="91"/>
</segment>
</net>
<net name="N$27" class="0">
<segment>
<pinref part="X8" gate="G$1" pin="B+"/>
<pinref part="X7" gate="G$1" pin="B+"/>
<wire x1="111.76" y1="30.48" x2="111.76" y2="35.56" width="0.1524" layer="91"/>
</segment>
</net>
<net name="N$28" class="0">
<segment>
<pinref part="X8" gate="G$1" pin="B-"/>
<pinref part="X7" gate="G$1" pin="C-"/>
<wire x1="114.3" y1="30.48" x2="119.38" y2="35.56" width="0.1524" layer="91"/>
</segment>
</net>
<net name="N$29" class="0">
<segment>
<pinref part="X8" gate="G$1" pin="C+"/>
<pinref part="X7" gate="G$1" pin="B-"/>
<wire x1="116.84" y1="30.48" x2="116.84" y2="33.02" width="0.1524" layer="91"/>
<wire x1="116.84" y1="33.02" x2="114.3" y2="35.56" width="0.1524" layer="91"/>
</segment>
</net>
<net name="N$30" class="0">
<segment>
<pinref part="X8" gate="G$1" pin="C-"/>
<pinref part="X7" gate="G$1" pin="C+"/>
<wire x1="119.38" y1="30.48" x2="119.38" y2="33.02" width="0.1524" layer="91"/>
<wire x1="119.38" y1="33.02" x2="116.84" y2="35.56" width="0.1524" layer="91"/>
</segment>
</net>
<net name="N$31" class="0">
<segment>
<pinref part="X8" gate="G$1" pin="D+"/>
<pinref part="X7" gate="G$1" pin="D+"/>
<wire x1="121.92" y1="30.48" x2="121.92" y2="35.56" width="0.1524" layer="91"/>
</segment>
</net>
<net name="N$32" class="0">
<segment>
<pinref part="X8" gate="G$1" pin="D-"/>
<pinref part="X7" gate="G$1" pin="D-"/>
<wire x1="124.46" y1="30.48" x2="124.46" y2="35.56" width="0.1524" layer="91"/>
</segment>
</net>
<net name="N$33" class="0">
<segment>
<pinref part="X10" gate="G$1" pin="A+"/>
<pinref part="X9" gate="G$1" pin="A+"/>
<wire x1="132.08" y1="30.48" x2="132.08" y2="35.56" width="0.1524" layer="91"/>
</segment>
</net>
<net name="N$34" class="0">
<segment>
<pinref part="X10" gate="G$1" pin="A-"/>
<pinref part="X9" gate="G$1" pin="A-"/>
<wire x1="134.62" y1="30.48" x2="134.62" y2="35.56" width="0.1524" layer="91"/>
</segment>
</net>
<net name="N$35" class="0">
<segment>
<pinref part="X10" gate="G$1" pin="B+"/>
<pinref part="X9" gate="G$1" pin="B+"/>
<wire x1="137.16" y1="30.48" x2="137.16" y2="35.56" width="0.1524" layer="91"/>
</segment>
</net>
<net name="N$36" class="0">
<segment>
<pinref part="X10" gate="G$1" pin="B-"/>
<pinref part="X9" gate="G$1" pin="C-"/>
<wire x1="139.7" y1="30.48" x2="144.78" y2="35.56" width="0.1524" layer="91"/>
</segment>
</net>
<net name="N$37" class="0">
<segment>
<pinref part="X10" gate="G$1" pin="C+"/>
<pinref part="X9" gate="G$1" pin="B-"/>
<wire x1="142.24" y1="30.48" x2="142.24" y2="33.02" width="0.1524" layer="91"/>
<wire x1="142.24" y1="33.02" x2="139.7" y2="35.56" width="0.1524" layer="91"/>
</segment>
</net>
<net name="N$38" class="0">
<segment>
<pinref part="X10" gate="G$1" pin="C-"/>
<pinref part="X9" gate="G$1" pin="C+"/>
<wire x1="144.78" y1="30.48" x2="144.78" y2="33.02" width="0.1524" layer="91"/>
<wire x1="144.78" y1="33.02" x2="142.24" y2="35.56" width="0.1524" layer="91"/>
</segment>
</net>
<net name="N$39" class="0">
<segment>
<pinref part="X10" gate="G$1" pin="D+"/>
<pinref part="X9" gate="G$1" pin="D+"/>
<wire x1="147.32" y1="30.48" x2="147.32" y2="35.56" width="0.1524" layer="91"/>
</segment>
</net>
<net name="N$40" class="0">
<segment>
<pinref part="X10" gate="G$1" pin="D-"/>
<pinref part="X9" gate="G$1" pin="D-"/>
<wire x1="149.86" y1="30.48" x2="149.86" y2="35.56" width="0.1524" layer="91"/>
</segment>
</net>
</nets>
</sheet>
</sheets>
</schematic>
</drawing>
</eagle>
