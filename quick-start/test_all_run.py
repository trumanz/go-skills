import os
import unittest
import subprocess


class Test(unittest.TestCase):
    def test_all_run_sh(self):
        #print("Current working dir : " + os.getcwd())
        cfd = os.path.dirname(__file__)
        #print("Current file dir : " +  cfd)
        i = 0
        for root,dirs,files in os.walk(cfd):
            for file in files:
                if file.startswith("run") and file.endswith(".sh"):
                    sh_file = root + "/" + file
                    i = i + 1
                    print("{}. Start Run {}".format(i,sh_file))
                    cp = subprocess.run(["bash", sh_file], cwd = root)
                    self.assertEqual(cp.returncode , 0)
        #cp = subprocess.run(["bash","run.sh"])
        #self.assertEqual(cp.returncode, 0)
