"""setup module"""

import setuptools

setuptools.setup(name='module_a',
                 version='1.1.32',
                 description='Module A',
                 url='https://github.com/sky-uk/ita-data-analytics',
                 author='sky',
                 zip_safe=False,
                 long_description_content_type='text/markdown',
                 packages=setuptools.find_packages(),
                 include_package_data=True,
                 install_requires=[
                     'mockito==1.1.1',
                     'google-cloud-storage==1.17.0',
                     'module_b',
                 ],
                 test_suite='module_a.tests',
                 scripts=['bin/run_module_a.py'])
